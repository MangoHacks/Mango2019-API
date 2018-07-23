// Package routes deals with the execution of requests
// to different routes.
package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MangoHacks/Mango2019-API/database"
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
	// Mimic the request body.
	type preregisterRequest struct {
		Email string `json:"email"`
	}
	// Declare a struct to fill with the request body.
	var prr preregisterRequest
	if err := web.ReadJSONBodyIntoStruct(r.Body, &prr); err != nil {
		if err := web.SendHTTPResponse(w, web.BadRequestError); err != nil {
			log.Fatal(err)
		}
		return
	}
	eml := prr.Email

	if err := db.InsertPreregistration(eml); err != nil {
		if err != database.ErrDuplicate {
			err = &web.InternalServerError
		}
		if err := web.SendHTTPResponse(w, err); err != nil {
			log.Fatal(err)
		}
	}
	if err := web.SendHTTPResponse(w, "user successfully preregistered"); err != nil {
		log.Fatal(err)
	}
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
func GetPreregistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type preregistrations []struct {
		Email     string    `json:"email"`
		Timestamp time.Time `json:"timestamp"`
	}

	q := `SELECT * FROM preregistrations
		ORDER BY timestamp ASC`
	rws, err := db.Query(q)
	if err != nil {
		if err := web.SendHTTPResponse(w, web.InternalServerError); err != nil {
			log.Fatal(err)
		}
		log.Printf("encountered error while selecting from preregistrations: %s", err.Error())
		return
	}

	var prrs preregistrations
	for rws.Next() {
		var eml string
		var t time.Time
		rws.Scan(&eml, &t)
		prrs = append(prrs, struct {
			Email     string    `json:"email"`
			Timestamp time.Time `json:"timestamp"`
		}{
			Email:     eml,
			Timestamp: t,
		})
	}
	b, err := json.Marshal(prrs)
	if err != nil {
		log.Printf("encountered error while retrieving preregistrations: %s" + err.Error())
		return
	}
	if err := web.SendHTTPResponse(w, b); err != nil {
		log.Fatal(err)
	}
}

// DeletePreregistration handles a DELETE request to /preregistration.
// We delete the row containing the email received in the JSON body from the
// database.
//
// We expect a JSON object in the request body, of form:
//  {
//  	"email": "example@google.com"
//  }
func DeletePreregistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type preregisterRequest struct {
		Email string `json:"email"`
	}

	var prr preregisterRequest
	if err := web.ReadJSONBodyIntoStruct(r.Body, &prr); err != nil {
		if err := web.SendHTTPResponse(w, web.BadRequestError); err != nil {
			log.Fatal(err)
		}
		return
	}
	eml := prr.Email
	q := `DELETE FROM preregistrations
	WHERE email = $1`
	if _, err := db.Exec(q, eml); err != nil {
		if err := web.SendHTTPResponse(w, web.InternalServerError); err != nil {
			log.Fatal(err)
		}
		return
	}
	if err := web.SendHTTPResponse(w, "OK"); err != nil {
		log.Fatal(err)
	}
}
