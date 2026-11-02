package integration

import (
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"testing"
	"time"
)

func waitForCalls(
	t *testing.T,
	d *mock_client.Client,
	count int,
) {
	t.Helper()
	deadline := time.Now().Add(10 * time.Second)

	for time.Now().Before(deadline) {
		if d.Calls() >= count {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf("timed out waiting for %d downstream triggers", count)
}
