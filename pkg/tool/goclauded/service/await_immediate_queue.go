package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"
	"time"
)

func (s *Service) AwaitImmediateQueue(
	x context.Context,
	sessionIdentifier string,
	callsign string,
	hold time.Duration,
) ([]queue.Entry, error) {
	c := s.notifier.Subscribe()
	defer s.notifier.Unsubscribe(c)
	t := time.NewTimer(hold)
	defer t.Stop()

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
		case <-t.C:
			return nil, nil
		case <-x.Done():
			return nil, nil
		case <-s.done:
			return nil, nil
		}
	}
}
