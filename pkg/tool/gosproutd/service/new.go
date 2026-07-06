package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
) *Service {
	return &Service{
		store:    s,
		notifier: n,
	}
}
