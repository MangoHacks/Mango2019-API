package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// preRegistration handles serves and handles requests to the preregistration page.
func preRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // If the client is sending us a payload (we define the method of an action in the template)
		fmt.Println(r.FormValue("email")) // Print the email on the server side.
		// TODO: Rather than simply print the email,
		// we need to store the user as a tuple in PostgreSQL.
	}
	t, err := template.ParseFiles("templates/prereg.html") // Prepare the template for serving via the http.ResponseWriter
	if err != nil {                                        // If an error was generated, we log fatally (equivalent to fmt.Println(err); os.Exit(1))
		log.Fatal(err)
	}
	t.Execute(w, nil) // Give the template to http.ResponseWriter to display to the user.
}

func main() {
	http.HandleFunc("/", preRegistration)        // Set a function to route a visit to the root directory to.
	log.Fatal(http.ListenAndServe(":8000", nil)) // Listen and serve will run until it hits an error and logs fatally.
}
