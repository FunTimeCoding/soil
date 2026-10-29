package store

import "database/sql"

type Store struct {
	database *sql.DB
	model    string
}
