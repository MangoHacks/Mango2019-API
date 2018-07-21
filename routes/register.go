// Package routes deals with the execution of requests
// to different routes.
package routes

import (
	"database/sql"
	"net/http"
)

// PostRegister handles a POST to /register.
func PostRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

// GetRegister handles a GET to /register.
func GetRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}
