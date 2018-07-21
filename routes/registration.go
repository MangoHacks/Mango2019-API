package routes

import (
	"database/sql"
	"net/http"
)

// PostRegistration handles a POST to /registration.
// We add the user information received as a row in the database.
//
// We expect a JSON body in the POST, of form:
//  {
//  	"email": "example@google.com",
//  	"...": "..."
//  }
func PostRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

// GetRegistration handles a GET to /registration.
// We retrieve all the registered users from the database.
//
// We send a JSON body in the response, of form:
//  [
//  	{
//  		"email": "example1@google.com",
//  		"...": "..."
//  	},
//  	{
//  		"email": "example2@google.com",
//  		"...": "..."
//  	}
//  ]
//
func GetRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}

// DeleteRegistration handles a DELETE request to /registration.
// We delete the row containing the email received in the JSON body from the
// database.
//
// We expect a JSON object in the request body, of form:
//  {
//  	"email": "example@google.com"
//  }
func DeleteRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}
