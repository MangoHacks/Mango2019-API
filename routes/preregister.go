package routes

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/MangoHacks/Mango2019-API/web"
)

// PostPreregister handles a POST request to /preregister.
func PostPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

	// Insert into PostgresSQL
	q := `INSERT INTO preregistrations(email, timestamp) 
	VALUES($1, $2) 
	RETURNING email`
	if _, err := db.Exec(q, eml, time.Now()); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			web.SendHTTPResponse(w, errors.New("user already exists"))
		} else {
			if err := web.SendHTTPResponse(w, web.InternalServerError); err != nil {
				log.Fatal(err)
			}
			log.Printf("encountered error while inserting into preregistrations: %s", err.Error())
		}
		return
	}
	if err := web.SendHTTPResponse(w, "user successfully preregistered"); err != nil {
		log.Fatal(err)
	}
}

// GetPreregister handles a GET request to /preregister.
func GetPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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

// DeletePreregister deletes a row matching the email given in a request from the database.
func DeletePreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
