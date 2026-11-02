package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
) *Service {
	return &Service{notifier: n, store: s}
}
