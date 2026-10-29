package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	mutex    sync.Mutex
	database *gorm.DB
	store    *store.Store
	notifier face.EventNotifier
}
