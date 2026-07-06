package service

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"

func New(s *store.Store) *Service {
	return &Service{store: s}
}
