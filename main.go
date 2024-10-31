package main

import (
	"github.com/rs/zerolog/log"

	"github.com/adrian-feijo-fagundes/my-api-golang/api"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()
	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Error to start server")
	}
}
