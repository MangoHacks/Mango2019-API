// Package routes deals with the execution of requests
// to different routes.
package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/database"
	"github.com/MangoHacks/Mango2019-API/models"
	"github.com/MangoHacks/Mango2019-API/web"
)

// PostPreregistration handles a POST request to /preregistration.
// We add the email received in the JSON body as a row in the database.
//
// We expect a JSON object in the request body, of form:
//  {
//  	"email": "example@google.com"
//  }
func PostPreregistration(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var prr models.Preregistration
	if err := web.ReadJSONBodyIntoStruct(r.Body, &prr); err != nil {
		web.SendHTTPResponse(w, err)
		return
	}

	if err := db.InsertPreregistration(prr.Email); err != nil {
		if err != database.ErrDuplicate {
			err = &web.ErrInternalServer
		}
		web.SendHTTPResponse(w, err)
		return
	}
	web.SendHTTPResponse(w, "user successfully preregistered")
}

// GetPreregistration handles a GET request to /preregistration.
// We retrieve the emails of every user to have preregistered from
// database.
//
// We send a JSON body in the response, of form:
//  [
//  	{
//  		"email" "example1@google.com"
//  	},
//  	{
//  		"email": "example2@google.com"
//  	}
//  ]
func GetPreregistration(w http.ResponseWriter, r *http.Request, db *database.DB) {
	prrs, err := db.SelectPreregistrations()
	if err != nil {
		web.SendHTTPResponse(w, web.ErrInternalServer)
		log.Printf("encountered error while retrieving preregistrations: %v", err)
		return
	}
	b, err := json.Marshal(prrs)
	if err != nil {
		log.Printf("encountered error while marshalling preregistrations: %v", err)
		return
	}
	web.SendHTTPResponse(w, b)
}

// DeletePreregistration handles a DELETE request to /preregistration.
// We delete the row containing the email received in the JSON body from the
// database.
//
// We expect a JSON object in the request body, of form:
//  {
//  	"email": "example@google.com"
//  }
func DeletePreregistration(w http.ResponseWriter, r *http.Request, db *database.DB) {
	var prr models.Preregistration
	if err := web.ReadJSONBodyIntoStruct(r.Body, &prr); err != nil {
		web.SendHTTPResponse(w, web.ErrBadRequest)
		log.Printf("encountered error parsing delete request into struct: %v", err)
		return
	}
	if err := db.DeletePreregistration(prr.Email); err != nil {
		web.SendHTTPResponse(w, web.ErrInternalServer)
		return
	}
	web.SendHTTPResponse(w, "OK")
}
