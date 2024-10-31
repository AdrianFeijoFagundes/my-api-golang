package main

import (
	"log"

	"github.com/adrian-feijo-fagundes/my-api-golang/api"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
