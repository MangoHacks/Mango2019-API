package main

import (
	"log"

	"github.com/MangoHacks/Mango2019-API/server"
)

func main() {
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
