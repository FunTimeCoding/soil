package service

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	database  *gorm.DB
	store     *store.Store
	notifier  *notifier.Notifier
	closeOnce sync.Once
}
