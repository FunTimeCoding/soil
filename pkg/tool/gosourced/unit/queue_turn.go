package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"testing"
)

func queueTurn(
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
	waitForQueue(t, s, queued+1)
}
