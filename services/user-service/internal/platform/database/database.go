package database

import (
	"context"
	"database/sql"
)


type DB interface {

	Close() error

	PingContext(
		ctx context.Context,
	) error


	QueryRowContext(
		ctx context.Context,
		query string,
		args ...any,
	) *sql.Row

	ExecContext(
		ctx context.Context,
		query string,
		args ...any,
	) (sql.Result,error)
}