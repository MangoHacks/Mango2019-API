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
		switch r.Method {
		case "POST":
			routes.PostPreregistration(w, r, db)
		case "GET":
			routes.GetPreregistration(w, r, db)
		case "DELETE":
			routes.DeletePreregistration(w, r, db)
		default:
			web.SendHTTPResponse(w, web.ErrMethodNotAllowed)
		}
	}
}

// handleRegister handles a request to /register and sends them to the appropriate route.
func handleRegistration(db *database.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			routes.PostRegistration(w, r, db)
		case "GET":
			routes.GetRegistration(w, r, db)
		case "DELETE":
			routes.DeleteRegistration(w, r, db)
		default:
			web.SendHTTPResponse(w, web.ErrMethodNotAllowed)
		}
	}
}
