package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(
	cfg config.DatabaseConfig,
) (DB,error){


	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := sql.Open(
		"pgx",
		dsn,
	)

	if err != nil {
		return nil,err
	}

	ctx,cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {

		return nil,err

	}

	return &Postgres{
		db: db,
	},nil
}



func (p *Postgres) Close() error {
	return p.db.Close()
}


func (p *Postgres) PingContext(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *Postgres) QueryRowContext(
	ctx context.Context,
	query string,
	args ...any,
) *sql.Row {

	return p.db.QueryRowContext(
		ctx,
		query,
		args...,
	)
}

func (p *Postgres) ExecContext(
	ctx context.Context,
	query string,
	args ...any,
) (sql.Result,error){

	return p.db.ExecContext(
		ctx,
		query,
		args...,
	)
}