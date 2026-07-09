package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(entry.New()))

	return &Store{database: m}
}
