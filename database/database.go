// Package database deals with the initialization
// of our database connection.
package database

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

var (
	// ErrDuplicate refers to an error generated when
	// a unique constraint is violated.
	ErrDuplicate = errors.New("database: duplicate entry")
)

// DB wraps the underlying database connection, allowing
// for methods to be executed through it.
type DB struct {
	postgres *sql.DB
}

// InsertPreregistration inserts a new preregistration into the appropriate
// database.
func (db *DB) InsertPreregistration(email string) error {
	if db.postgres != nil {
		return db.postgresInsertPreregistration(email)
	}
	return nil
}

func (db *DB) postgresInsertPreregistration(email string) error {
	q := `INSERT INTO preregistrations(email, timestamp) 
	VALUES($1, $2) 
	RETURNING email`
	if _, err := db.postgres.Exec(q, email, time.Now()); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return ErrDuplicate
		}
		return err
	}
	return nil
}
