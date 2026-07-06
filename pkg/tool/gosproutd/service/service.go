package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"sync"
)

type Service struct {
	mutex    sync.Mutex
	store    *store.Store
	notifier face.EventNotifier
}
