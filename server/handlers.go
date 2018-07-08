package server

import (
	"database/sql"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/routes"
)

// handlePreregister handles a request to /preregister and sends them to the appropriate route.
func handlePreregister(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			routes.PostPreregister(w, r, db)
		} else if r.Method == "GET" {
			routes.GetPreregister(w, r, db)
		}
	}
}

// handleRegister handles a request to /register and sends them to the appropriate route.
func handleRegister(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			routes.PostRegister(w, r, db)
		} else if r.Method == "GET" {
			routes.GetRegister(w, r, db)
		}
	}
}