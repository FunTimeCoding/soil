package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(label.Stub()))

	return &Store{database: m}
}
