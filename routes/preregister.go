package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// PostPreregister handles a POST request to /preregister.
func PostPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type PreregisterRequest struct {
		Email string `json:"email"`
	}
	var prr PreregisterRequest

	if err := readJSONBodyIntoStruct(r.Body, &prr); err != nil {
		sendJSONResponse(w, 400, []byte(err.Error()))
		return
	}
	eml := prr.Email

	var e string
	if err := db.QueryRow(`INSERT INTO preregistrations(email) 
		VALUES($1) 
		RETURNING email`, eml).Scan(&e); err != nil {
		sendJSONResponse(w, 400, []byte(err.Error()))
		return
	}
	sendJSONResponse(w, 200, []byte("OK"))
}

// GetPreregister handles a GET request to /preregister.
func GetPreregister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	type Preregistrations []struct {
		Email string `json:"email"`
	}

	rws, err := db.Query("SELECT * FROM preregistrations")
	if err != nil {
		sendJSONResponse(w, 400, []byte(err.Error()))
		return
	}

	var prrs Preregistrations
	for rws.Next() {
		var eml string
		rws.Scan(&eml)
		prrs = append(prrs, struct {
			Email string `json:"email"`
		}{
			Email: eml,
		})
	}
	b, err := json.Marshal(prrs)
	if err != nil {
		sendJSONResponse(w, 400, []byte(err.Error()))
		return
	}
	sendJSONResponse(w, 200, b)
}
