package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func New(
	s *store.Store,
	q face.Client,
) *Service {
	return &Service{store: s, query: q}
}
