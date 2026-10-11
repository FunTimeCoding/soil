package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) pushEntry(e *queue.Entry) error {
	if f := s.store.PushEntry(e); f != nil {
		return f
	}

	s.notify()

	return nil
}
