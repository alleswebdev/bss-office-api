package database

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-office-api/internal/logger"
)

// StatementBuilder глобальная переменная с сконфигурированным плейсхолдером для pgsql
var StatementBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

// NewPostgres returns DB
func NewPostgres(ctx context.Context, dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		logger.ErrorKV(ctx, "failed to create database connection", "err", err)

		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.ErrorKV(ctx, "failed ping the database", "err", err)

		return nil, err
	}

	return db, nil
}
