package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(path string) *Store {
	d, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(d.Exec("PRAGMA journal_mode = WAL").Error)
	errors.PanicOnError(d.AutoMigrate(seed.Stub()))

	return &Store{mapper: d}
}
