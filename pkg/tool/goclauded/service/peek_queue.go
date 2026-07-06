package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) PeekQueue(callsign string) ([]queue.Entry, error) {
	return s.store.PeekQueue(callsign)
}
