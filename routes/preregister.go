package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// PostPreregister handles a POST request to /preregister.
func PostPreregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.FormValue("email"))
}

// GetPreregister handles a GET request to /preregister.
func GetPreregister(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/prereg.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
