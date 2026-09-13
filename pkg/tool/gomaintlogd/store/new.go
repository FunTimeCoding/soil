package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	n face.EventNotifier,
) *Store {
	errors.PanicOnError(m.AutoMigrate(entry.New()))

	return &Store{database: m, notifier: n}
}
