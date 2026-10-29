package store

import "database/sql"

func New(
	d *sql.DB,
	model string,
) *Store {
	initialize(d)

	return &Store{database: d, model: model}
}
