package pg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

// NewPG initializes the PostgreSQL connection pool with a singleton pattern.
func NewPG(ctx context.Context, connString string) (*postgres, error) {
	var initErr error // define a local error variable

	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			initErr = fmt.Errorf("unable to create connection pool: %w", err)
			return
		}
		pgInstance = &postgres{db: db}
	})

	// If initialization failed, return the error
	if initErr != nil {
		return nil, initErr
	}

	return pgInstance, nil
}

// GetDB exposes the pgxpool.Pool to allow query execution.
func (pg *postgres) GetDB() *pgxpool.Pool {
	if pg == nil || pg.db == nil {
		return nil // return nil if not initialized
	}
	return pg.db
}

// Ping checks if the database connection is alive.
func (pg *postgres) Ping(ctx context.Context) error {
	if pg == nil || pg.db == nil {
		return fmt.Errorf("database not initialized")
	}
	return pg.db.Ping(ctx)
}

// Close closes the database connection pool with context.
func (pg *postgres) Close(ctx context.Context) {
	if pg.db != nil {
		pg.db.Close()
	}
}
