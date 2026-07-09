package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(NewClick()))

	return &Store{mapper: m}
}
