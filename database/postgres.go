package database

import (
	"database/sql"
	"fmt"
	"os"

	// To avoid using pq directly.
	_ "github.com/lib/pq"
)

// New constructs a new database connection.
func New() (*sql.DB, error) {
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DBUser, DBPassword, DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
