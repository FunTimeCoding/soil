package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *Store {
	errors.PanicOnError(d.AutoMigrate(mute_rule.Stub()))

	return &Store{database: d}
}
