package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *Store {
	errors.PanicOnError(d.AutoMigrate(seed.Stub()))

	return &Store{mapper: d}
}
