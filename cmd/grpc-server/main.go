package main

import (
	"github.com/chillyNick/librarySearch/internal/config"
	"github.com/chillyNick/librarySearch/internal/database"
	"github.com/chillyNick/librarySearch/internal/repo/sqlx_repo"
	"github.com/chillyNick/librarySearch/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := config.ReadConfigYML("config.yaml"); err != nil {
		log.Fatal().Err(err).Msg("Failed to read configuration")
	}

	cfg := config.GetConfigInstance()

	db, err := database.NewMysql(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to init mysql")
	}

	if err := server.NewGrpcServer(sqlx_repo.New(db)).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}
