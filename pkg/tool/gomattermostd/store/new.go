package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *Store {
	errors.PanicOnError(d.AutoMigrate(subscription.Stub()))

	return &Store{database: d}
}
