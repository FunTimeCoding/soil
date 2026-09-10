package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) DrainQueue(
	sessionIdentifier string,
	callsign string,
) ([]queue.Entry, error) {
	return s.store.DrainQueue(sessionIdentifier, callsign)
}
