package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	n face.EventNotifier,
) *Store {
	return &Store{database: m, notifier: n}
}
