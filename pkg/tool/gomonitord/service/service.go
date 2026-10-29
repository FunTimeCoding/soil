package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"gorm.io/gorm"
)

type Service struct {
	database *gorm.DB
	notifier face.EventNotifier
	store    *store.Store
}
