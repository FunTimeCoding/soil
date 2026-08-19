package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store/record"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *Store {
	errors.PanicOnError(d.AutoMigrate(record.Stub()))

	return &Store{database: d}
}
