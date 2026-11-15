package turn_tester

import (
	"github.com/funtimecoding/soil/pkg/time/wait_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn"
	"testing"
	"time"
)

func Wait(
	t *testing.T,
	s *turn.Scheduler,
	count int,
) {
	t.Helper()
	wait_tester.Until(
		t,
		func() bool { return s.Waiting() >= count },
		5*time.Second,
	)
}
