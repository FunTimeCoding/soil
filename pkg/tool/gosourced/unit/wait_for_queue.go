package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/turn"
	"testing"
	"time"
)

func waitForQueue(
	t *testing.T,
	s *turn.Scheduler,
	count int,
) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)

	for s.Waiting() < count {
		if time.Now().After(deadline) {
			t.Fatalf("queue never reached %d", count)
		}

		time.Sleep(time.Millisecond)
	}
}
