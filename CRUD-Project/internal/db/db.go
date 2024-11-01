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

// NewPG initializes the PostgreSQL connection pool.
// It is safe to call multiple times, but it will only create a single instance.
func NewPG(ctx context.Context, connString string) (*postgres, error) {
	var err error

	pgOnce.Do(func() {
		db, errInit := pgxpool.New(ctx, connString)
		if errInit != nil {
			err = fmt.Errorf("unable to create connection pool: %w", errInit)
			return
		}

		pgInstance = &postgres{db}
	})

	if err != nil {
		return nil, err // Return the error if initialization failed
	}
	return pgInstance, nil
}

// Ping checks if the database connection is alive.
func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

// Close closes the database connection pool.
func (pg *postgres) Close() {
	if pg.db != nil {
		pg.db.Close()
	}
}
