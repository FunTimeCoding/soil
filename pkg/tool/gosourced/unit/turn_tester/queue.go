package turn_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"testing"
)

func Queue(
	t *testing.T,
	s *turn.Scheduler,
	session string,
	name string,
	granted chan<- string,
) {
	t.Helper()
	queued := s.Waiting()
	go func() {
		g, e := s.Take(context.Background(), session)

		if e != nil {
			granted <- e.Error()

			return
		}

		granted <- name
		g.Release()
	}()
	Wait(t, s, queued+1)
}
