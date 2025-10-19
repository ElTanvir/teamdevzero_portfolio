package main

import (
	"os"
	"portfolioed/internal/config"
	"portfolioed/internal/modules/menu"
	"portfolioed/internal/modules/root"
	"portfolioed/internal/server"
	"portfolioed/util"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := config.Load()
	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	util.RegisterTagName()
	app, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create server")
	}
	root.Init(app)
	menu.Init(app)
	log.Fatal().Err(app.Start()).Msg("failed to start server")
}
