package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func (s *Service) PushQueueImmediate(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) (bool, error) {
	return s.pushImmediate(queue.New(sessionIdentifier, callsign, kind, body))
}
