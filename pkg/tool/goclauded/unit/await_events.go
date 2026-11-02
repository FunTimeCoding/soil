package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	peerConstant "github.com/funtimecoding/soil/pkg/peer/constant"
	"testing"
	"time"
)

func awaitEvents(
	t *testing.T,
	r *memory.Memory,
	count int,
) {
	t.Helper()
	deadline := time.Now().Add(peerConstant.FixtureSignalWait)

	for len(r.Events()) < count && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
}
