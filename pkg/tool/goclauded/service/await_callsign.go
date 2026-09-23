package service

import (
	"context"
	"time"
)

func (s *Service) AwaitCallsign(
	x context.Context,
	sessionIdentifier string,
	since time.Time,
) (string, error) {
	c := s.notifier.Subscribe()
	defer s.notifier.Unsubscribe(c)

	for {
		callsign, e := s.store.LiveCallsign(sessionIdentifier, since)

		if e != nil {
			return "", e
		}

		if callsign != "" {
			return callsign, nil
		}

		select {
		case <-c:
		case <-x.Done():
			return "", nil
		case <-s.done:
			return "", nil
		}
	}
}
