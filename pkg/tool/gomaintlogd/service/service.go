package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
)

type Service struct {
	store    *store.Store
	notifier face.EventNotifier
}
