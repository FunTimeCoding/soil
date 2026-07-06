package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newLite(path string) *Store {
	database, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(NewClick()))

	return &Store{mapper: database}
}
