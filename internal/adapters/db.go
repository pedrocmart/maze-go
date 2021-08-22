package adapters

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func NewDBConnection(dbURL string, timeout, refresh int) (*sql.DB, error) {
	if timeout < 0 {
		return nil, errors.New("invalid timeout value provided")
	}

	if refresh < 0 {
		return nil, errors.New("invalid refresh value provided")
	}

	var db *sql.DB
	var err error

	connected := make(chan bool)

	go func() {
		for {
			// Connect to the database.
			db, err = sql.Open("postgres", "postgresql://maze:maze@maze:5432/maze_db?sslmode=disable&log_statement=all")
			if err == nil {
				// Ping the database.
				err = db.Ping()
				if err == nil {
					log.Println("Connected to PostgreSQL")
					connected <- true
					break
				}
			}

			// Sleep one second if connection failed.
			log.Printf("Failed to connect to PostgreSQL on %s, reconnecting in %d seconds", dbURL, refresh)
			time.Sleep(time.Duration(refresh) * time.Second)
		}
	}()

	select {
	case <-connected:
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil, fmt.Errorf("timeout occured connecting to PostgreSQL after %d seconds", timeout)
	}

	return db, nil
}

func ReturnInvalidMaxLengthError(f string, max int) error {
	errorMessage := fmt.Sprintf("%s max length is %v", f, max)
	return errors.New(errorMessage)
}

func ReturnEmptyFieldError(f string) error {
	errorMessage := fmt.Sprintf("%s cannot be empty", f)
	return errors.New(errorMessage)
}

func ReturnError(f string) error {
	return errors.New(f)
}
