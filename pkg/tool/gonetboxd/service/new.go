package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func New(
	c face.NetboxSource,
	s *store.Store,
) *Service {
	return &Service{client: c, store: s}
}
