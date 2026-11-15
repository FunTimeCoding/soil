package event_tester

import (
	"testing"
	"time"
)

func Next(
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
