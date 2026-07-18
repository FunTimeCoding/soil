package store

import (
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/channel"
	"github.com/funtimecoding/soil/pkg/chat/telegram/database/user"
	"github.com/funtimecoding/soil/pkg/errors"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(channel.Stub(), user.Stub()))

	return &Store{database: m}
}
