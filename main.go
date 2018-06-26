package main

import (
	"log"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/server"
)

func main() {
	s := server.New()
	if err := http.ListenAndServe(":8000", s.Router); err != nil {
		log.Fatal(err)
	}
}
