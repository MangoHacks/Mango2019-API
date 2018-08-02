package server

import (
	"net/http"

	"github.com/MangoHacks/Mango2019-API/database"
	"github.com/MangoHacks/Mango2019-API/web"

	"github.com/MangoHacks/Mango2019-API/routes"
)

// handlePreregister handles a request to /preregister and sends them to the appropriate route.
func handlePreregistration(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			routes.PostPreregistration(w, r, db)
		} else if r.Method == "GET" {
			routes.GetPreregistration(w, r, db)
		} else if r.Method == "DELETE" {
			routes.DeletePreregistration(w, r, db)
		} else {
			web.SendHTTPResponse(w, web.MethodNotAllowedError)
		}
	}
}

// handleRegister handles a request to /register and sends them to the appropriate route.
func handleRegistration(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			routes.PostRegistration(w, r, db)
		} else if r.Method == "GET" {
			routes.GetRegistration(w, r, db)
		} else if r.Method == "DELETE" {
			routes.DeleteRegistration(w, r, db)
		} else {
			web.SendHTTPResponse(w, web.MethodNotAllowedError)
		}
	}
}
