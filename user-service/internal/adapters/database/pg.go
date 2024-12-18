package database

import (
	"context"
	"enceland_user-service/internal/core/ports"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXAdapter struct {
	pool *pgxpool.Pool
}

func NewPGXAdapter(connStr string) (*PGXAdapter, error) {
	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	cfg.MinConns = 5
	cfg.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	if m, err := migrate.New("file://migrations", connStr); err != nil {
		return nil, err
	} else {
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	return &PGXAdapter{pool: pool}, nil
}

func (db *PGXAdapter) Query(ctx context.Context, query string, args ...any) (ports.DBRows, error) {
	rows, err := db.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
