package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newLite(path string) *Store {
	database, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(entry.New()))

	return &Store{database: database}
}
