package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) pushEntryBroadcast(template *queue.Entry) error {
	sessions, e := s.store.ListSessions()

	if e != nil {
		return e
	}

	if f := s.store.PushQueueBroadcast(sessions, template); f != nil {
		return f
	}

	s.notify()

	return nil
}
