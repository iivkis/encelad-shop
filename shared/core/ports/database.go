package ports

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrDB       = errors.New("database error")
	ErrNotFound = fmt.Errorf("%w: record not found", ErrDB)
)

type DBRows interface {
	Next() bool
	Values() ([]any, error)
	Close()
}

type DBConn interface {
	Query(ctx context.Context, query string, args ...any) (DBRows, error)
}
