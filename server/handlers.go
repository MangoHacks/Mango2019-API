package server

import (
	"net/http"

	"github.com/MangoHacks/Mango2019-API/routes"
)

// HandlePreregister routes a request to /preregister to the appropriate route.
func handlePreregister() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" { // If the client is sending us a payload (we define the method of an action in the template)
			routes.PostPreregister(w, r)
		} else if r.Method == "GET" {
			routes.GetPreregister(w, r)
		}
	}
}
