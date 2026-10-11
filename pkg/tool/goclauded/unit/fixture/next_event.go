package fixture

import (
	"testing"
	"time"
)

func NextEvent(
	t *testing.T,
	events <-chan Event,
) *Event {
	t.Helper()

	select {
	case e := <-events:
		return &e
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for SSE event")

		return nil
	}
}
