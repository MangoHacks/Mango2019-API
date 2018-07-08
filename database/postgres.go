package database

import (
	"database/sql"

	// To avoid using pq directly.
	_ "github.com/lib/pq"
)

// New constructs a new database connection.
func New() (*sql.DB, error) {
	connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
