package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"message_service/config"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), cfg.DatabaseURLA)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
