package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Service) GetEvent(identifier uint) *event.Event {
	return s.store.GetEvent(identifier)
}
