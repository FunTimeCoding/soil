package service

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Service) PushQueueImmediate(
	sessionIdentifier string,
	callsign string,
	kind string,
	body string,
) (bool, error) {
	recent, e := s.store.CountImmediateQueue(
		sessionIdentifier,
		callsign,
		s.clock().Add(-constant.ImmediateWindow),
	)

	if e != nil {
		return false, e
	}

	if recent >= constant.ImmediateLimit {
		return false, s.PushQueue(sessionIdentifier, callsign, kind, body)
	}

	if f := s.store.PushQueueImmediate(
		sessionIdentifier,
		callsign,
		kind,
		body,
	); f != nil {
		return false, f
	}

	s.notify()

	return true, nil
}
