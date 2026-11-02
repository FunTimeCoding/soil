package service

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

func New(
	s *store.Store,
	r face.Reporter,
) *Service {
	return &Service{store: s, reporter: r}
}
