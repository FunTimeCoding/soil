package turn_tester

import (
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"time"
)

func (o *Tester) Wait(
	s *turn.Scheduler,
	count int,
) {
	o.t.Helper()
	wait_tester.Until(
		o.t,
		func() bool { return s.Waiting() >= count },
		5*time.Second,
	)
}
