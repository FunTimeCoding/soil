package client

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"testing"
	"time"
)

func nextEvent(
	t *testing.T,
	events chan []client.Claim,
) []client.Claim {
	t.Helper()

	select {
	case v := <-events:
		return v
	case <-time.After(5 * time.Second):
		t.Fatal("no event within 5 seconds")

		return nil
	}
}
