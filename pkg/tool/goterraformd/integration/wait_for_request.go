package integration

import (
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"testing"
	"time"
)

func waitForRequest(
	t *testing.T,
	n *mock_runner.Runner,
) []trigger.Request {
	t.Helper()
	deadline := time.Now().Add(10 * time.Second)

	for time.Now().Before(deadline) {
		if r := n.Requests(); len(r) > 0 {
			return r
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatal("timed out waiting for the downstream trigger")

	return nil
}
