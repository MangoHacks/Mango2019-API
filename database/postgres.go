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

// Database Credentials
//
// These are the credentials necessary to initialize a
// connection with the database.
var (
	//////////////////////////////////
	// PostgreSQL Database Credentials
	//////////////////////////////////
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
)

// New returns a new connection to the specified PostgreSQL database.
//
// The credentials for the database are exepected to be exported and
// will be pulled down from the environment.
func New() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
