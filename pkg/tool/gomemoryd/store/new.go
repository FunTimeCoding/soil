package store

import "database/sql"

func New(d *sql.DB) *Store {
	initialize(d)

	return &Store{database: d}
}
