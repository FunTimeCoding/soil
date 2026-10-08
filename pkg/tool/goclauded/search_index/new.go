package search_index

import "database/sql"

func New(d *sql.DB) *Index {
	initialize(d)

	return &Index{database: d}
}
