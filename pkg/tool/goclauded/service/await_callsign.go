package service

import (
	"context"
	"time"
)

func (s *Service) AwaitCallsign(
	x context.Context,
	sessionIdentifier string,
	since time.Time,
	hold time.Duration,
) (string, error) {
	c := s.notifier.Subscribe()
	defer s.notifier.Unsubscribe(c)
	t := time.NewTimer(hold)
	defer t.Stop()

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
		case <-t.C:
			return "", nil
		case <-x.Done():
			return "", nil
		case <-s.done:
			return "", nil
		}
	}
}
