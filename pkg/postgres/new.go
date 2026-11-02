package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(locator string) (*Client, error) {
	p, e := pgxpool.New(context.Background(), locator)

	if e != nil {
		return nil, e
	}

	return &Client{pool: p}, nil
}
