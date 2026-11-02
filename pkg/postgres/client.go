package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type Client struct {
	pool *pgxpool.Pool
}
