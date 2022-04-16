package storage

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	// The postgresql connector
	_ "github.com/lib/pq"
)

const driver = "postgres"

func NewPostgres(ctx context.Context, dsn string) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return db, nil
}
