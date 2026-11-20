package turn_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
)

func (o *Tester) Queue(
	s *turn.Scheduler,
	session string,
	name string,
	granted chan<- string,
) {
	o.t.Helper()
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
	o.Wait(s, queued+1)
}
