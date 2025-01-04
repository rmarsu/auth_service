package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	return dbpool, nil
}
