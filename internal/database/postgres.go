package database

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

// StatementBuilder глобальная переменная с установленным долларом в формте pgsql
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// NewPostgres returns DB
func NewPostgres(dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed ping the database")

		return nil, err
	}

	return db, nil
}
