package store

import (
	"database/sql"
	"github.com/funtimecoding/soil/pkg/errors"
	_ "github.com/glebarez/go-sqlite"
)

func New(path string) *Store {
	database, e := sql.Open("sqlite", path)
	errors.PanicOnError(e)
	initialize(database)

	return &Store{database: database}
}
