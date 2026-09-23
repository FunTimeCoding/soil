package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
)

func (s *Service) AwaitImmediateQueue(
	x context.Context,
	sessionIdentifier string,
	callsign string,
) ([]queue.Entry, error) {
	c := s.notifier.Subscribe()
	defer s.notifier.Unsubscribe(c)

	for {
		drained, e := s.DrainImmediateQueue(sessionIdentifier, callsign)

		if e != nil {
			return nil, e
		}

		if len(drained) > 0 {
			return drained, nil
		}

		select {
		case <-c:
		case <-x.Done():
			return nil, nil
		case <-s.done:
			return nil, nil
		}
	}
}
