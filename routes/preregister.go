package routes

import (
	"database/sql"
	"net/http"
)

// PostPreregister handles a POST request to /preregister.
func PostPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: Make query to PostgresSQL to add preregistration.
}

// GetPreregister handles a GET request to /preregister.
func GetPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: Make query to PostgresSQL to read preregistrations.
}
