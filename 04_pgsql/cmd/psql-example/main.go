package main

import (
	"os"

	"github.com/itiky/practicum-examples/04_pgsql/cmd/psql-example/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Application crashed")
		os.Exit(1)
	}
}
