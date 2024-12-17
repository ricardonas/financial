package database

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

// NewPG initializes the PostgreSQL connection and returns it
func NewPG(ctx context.Context, connString string) (*postgres, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			pgInstance = nil
			return
		}
		pgInstance = &postgres{db}
	})
	if pgInstance == nil {
		return nil, fmt.Errorf("unable to create connection pool")
	}
	return pgInstance, nil
}

// Getter method to access the pgxpool.Pool
func (pg *postgres) GetDB() *pgxpool.Pool {
	return pg.db
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}
