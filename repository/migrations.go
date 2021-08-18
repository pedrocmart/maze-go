package repository

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/gobuffalo/packr"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

type Migrations interface {
	RunMigrations() error
}

type MigrationsRepository struct {
	db            *sql.DB
	migrationsBox packr.Box
}

func NewMigrationsRepository(db *sql.DB, migrationsBox packr.Box) *MigrationsRepository {
	return &MigrationsRepository{
		db:            db,
		migrationsBox: migrationsBox,
	}
}

func (r *MigrationsRepository) PrintMigrations() error {
	for _, filename := range r.migrationsBox.List() {
		fmt.Printf("migration %s\n", filename)

		f, fErr := r.migrationsBox.Open(filename)
		if fErr != nil {
			return fmt.Errorf("failed to open migration file %s: %v", filename, fErr)
		}
		io.Copy(os.Stdout, f)
		fmt.Println()
	}

	return nil
}

func (r *MigrationsRepository) RunMigrations() (int, error) {
	migrations := &migrate.PackrMigrationSource{
		Box: r.migrationsBox,
	}

	migrate.SetTable("migrations")

	n, err := migrate.Exec(r.db, "postgres", migrations, migrate.Up)
	return n, err
}

func (r *MigrationsRepository) RollbackMigration() (int, error) {
	migrations := &migrate.PackrMigrationSource{
		Box: r.migrationsBox,
	}

	migrate.SetTable("migrations")

	n, err := migrate.Exec(r.db, "postgres", migrations, migrate.Down)
	return n, err
}
