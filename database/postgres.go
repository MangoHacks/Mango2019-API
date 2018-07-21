// Package database deals with the initialization
// of our database connection.
package database

import (
	"database/sql"
	"fmt"
	"os"

	// To avoid using pq directly.
	_ "github.com/lib/pq"
)

// Database credentials
//
// These are the credentials necessary to initialize a
// connection with th database.
var (
	// DBUser is the username for the database.
	DBUser string

	// DBPassword is the password for the database.
	DBPassword string

	// DBName is the name of the database.
	DBName string
)

// New returns a new connection to the specified database.
//
// The credentials for the database are exepected to be exported and
// will be pulled down from the environment.
func New() (*sql.DB, error) {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
