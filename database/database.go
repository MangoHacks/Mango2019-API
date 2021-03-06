// Package database deals with the initialization
// of our database connection.
package database

import (
	"database/sql"
	"errors"
	"log"

	"github.com/MangoHacks/Mango2019-API/models"
)

// Database Errors
//
// These are the errors that can be generated by our databases.
var (
	// ErrDuplicate refers to an error generated when
	// a unique constraint is violated.
	ErrDuplicate = errors.New("database: duplicate entry")
)

// DB wraps the underlying database connection, allowing
// for methods to be executed through it.
//
// The purpose of this abstraction is to allow for the use of different
// databases without forcing the user to undergo a major refactor.
//
// For the time being, planned support is for the following databases:
// * PostgreSQL (Currently Supported / Used by MangoHacks 2019)
// * MongoDB
// * Google Firebase
type DB struct {
	postgres *sql.DB
}

// New returns a new connection to the specified PostgreSQL database.
//
// The credentials for the database are exepected to be exported and
// will be pulled down from the environment.
func New(database string) (*DB, error) {
	if database == "postgres" {
		db, err := newPostgres()
		if err != nil {
			return nil, err
		}
		log.Println("Successfully connected.")
		return &DB{
			postgres: db,
		}, nil
	}
	return nil, nil
}

// InsertPreregistration inserts a new preregistration into the appropriate
// database.
func (db *DB) InsertPreregistration(email string) error {
	if db.postgres != nil {
		return db.postgresInsertPreregistration(email)
	}
	return nil
}

// SelectPreregistrations selects all the preregistrations from the appropriate database.
func (db *DB) SelectPreregistrations() ([]models.Preregistration, error) {
	var prrs []models.Preregistration
	var err error
	if db.postgres != nil {
		prrs, err = db.postgresSelectPreregistrations()
		if err != nil {
			return nil, err
		}
	}
	return prrs, nil
}

// DeletePreregistration deletes the given preregistration from the appropriate database.
func (db *DB) DeletePreregistration(email string) error {
	if db.postgres != nil {
		return db.postgresDeletePreregistration(email)
	}
	return nil
}
