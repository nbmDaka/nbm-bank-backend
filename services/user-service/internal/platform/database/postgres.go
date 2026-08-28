package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgres(cfg config.DatabaseConfig) (*sql.DB, error) {

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
		return nil, err
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
