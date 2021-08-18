package main

import (
	"database/sql"
	"os"
	"sync"

	"github.com/pedrocmart/maze-go/api/restapi"
	"github.com/pedrocmart/maze-go/api/restapi/operations"
	"github.com/pedrocmart/maze-go/consts"
	"github.com/pedrocmart/maze-go/handlers"
	"github.com/pedrocmart/maze-go/internal/adapters"
	"github.com/pedrocmart/maze-go/internal/config"
	"github.com/pedrocmart/maze-go/internal/logger"
	"github.com/pedrocmart/maze-go/repository"

	"github.com/go-openapi/loads"
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
)

var loggerClient log.FieldLogger

type operatorAPIs struct {
	mu   sync.Mutex
	apis map[int64]*opClient.OperatorAPIRequirements
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configs: %v", err)
	}

	loggerClient = logger.NewLogger(cfg.LogOutput, cfg.LogFormat)

	// Establish DB connection.
	db, err := adapters.NewDBConnection(cfg.DBURL, cfg.DBTimeout, cfg.DBRefresh)
	if err != nil {
		loggerClient.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to establish DB connection")
	}

	loggerClient.Info("Established DB connection")

	// Handle migrations.
	handleMigrations(cfg, db)

	// Load embedded swagger file.
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		loggerClient.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to load server swagger file")
	}

	// Create new service API.
	api := operations.NewMazeGoAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	api.Logger = loggerClient.Printf

	server.Host = cfg.Host
	server.Port = cfg.Port

	levelRepository := repository.NewLevelRepository(db)

	handlers.RegisterLevelHandlers(api, levelRepository)

	// Serve API.
	if err := server.Serve(); err != nil {
		loggerClient.WithFields(log.Fields{
			"error": err,
		}).Fatal("failed to serve API")
	}
}

func handleMigrations(cfg *config.Config, db *sql.DB) {
	var mRun int
	var err error
	migrationsRepo := repository.NewMigrationsRepository(db, packr.NewBox("../migrations"))

	switch {
	case cfg.Migrate == consts.MigratePrint:
		err = migrationsRepo.PrintMigrations()
		if err != nil {
			loggerClient.WithFields(log.Fields{
				"error": err,
			}).Fatal("failed to print database migrations")
		}

	case cfg.Migrate == "" || cfg.Migrate == consts.MigrateUp:
		loggerClient.Info("Running DB migrations")

		mRun, err = migrationsRepo.RunMigrations()
		if err != nil {
			loggerClient.WithFields(log.Fields{
				"error": err,
			}).Fatal("failed to run database migrations")
		}

	case cfg.Migrate == consts.MigrateDown:
		loggerClient.Info("Rolling back DB migrations")

		mRun, err = migrationsRepo.RollbackMigration()
		if err != nil {
			loggerClient.WithFields(log.Fields{
				"error": err,
			}).Fatal("failed to roll back database migration")
		}
	}

	loggerClient.WithField("migrations", mRun).Infof("executed migrations")

	// Terminate the process if only migration is required.
	if cfg.Migrate == consts.MigrateUp || cfg.Migrate == consts.MigrateDown || cfg.Migrate == consts.MigratePrint {
		loggerClient.Info("Terminating the process after running the migrations")
		os.Exit(0)
	}
}
