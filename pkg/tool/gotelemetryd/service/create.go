package service

import "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"

func (s *Service) Create(e *usage_event.Event) error {
	return s.store.Create(e)
}
