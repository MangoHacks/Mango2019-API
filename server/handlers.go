// Package server deals with the initialization of the server as
// well as the listening and handling of resources.
package server

import (
	"database/sql"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/web"

	"github.com/MangoHacks/Mango2019-API/routes"
)

// handlePreregister handles a request to /preregister and sends them to the appropriate route.
func handlePreregister(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			routes.PostPreregister(w, r, db)
		} else if r.Method == "GET" {
			routes.GetPreregister(w, r, db)
		} else if r.Method == "DELETE" {
			routes.DeletePreregister(w, r, db)
		} else {
			web.SendHTTPResponse(w, web.MethodNotAllowedError)
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
		} else {
			web.SendHTTPResponse(w, web.MethodNotAllowedError)
		}
	}
}
