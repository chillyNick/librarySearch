package database

import (
	"fmt"
	"time"

	"github.com/chillyNick/librarySearch/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// NewMysql connect and ping mysql returns sqlx.DB instance
func NewMysql(cfg config.Database) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name),
	)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}

	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.ConnMaxLifetime))
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed to ping mysql")

		return nil, err
	}

	return db, nil
}
