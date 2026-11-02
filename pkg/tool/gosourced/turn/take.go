package turn

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/turn_grant"
	"time"
)

func (s *Scheduler) Take(
	x context.Context,
	session string,
) (*turn_grant.Grant, error) {
	start := time.Now()
	s.mutex.Lock()

	if !s.busy {
		s.busy = true
		s.mutex.Unlock()

		return turn_grant.New(0, 0, s.release), nil
	}

	ahead := s.queued() + 1
	w := &waiter{ready: make(chan struct{})}

	if len(s.waiting[session]) == 0 {
		s.order = append(s.order, session)
	}

	s.waiting[session] = append(s.waiting[session], w)
	s.mutex.Unlock()
	t := time.NewTimer(s.bound)
	defer t.Stop()

	select {
	case <-w.ready:
		return turn_grant.New(ahead, time.Since(start), s.release), nil
	case <-x.Done():
		s.abandon(session, w)

		return nil, x.Err()
	case <-t.C:
		s.abandon(session, w)

		return nil, conflict.Format("busy, %d calls ahead", ahead)
	}
}
