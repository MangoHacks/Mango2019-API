// Package main initializes and runs the server.
// This package is small, because it is main.
package main

import (
	"log"

	"github.com/MangoHacks/Mango2019-API/server"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
