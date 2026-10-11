package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) PushQueue(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) error {
	return s.pushEntry(queue.New(sessionIdentifier, callsign, kind, body))
}
