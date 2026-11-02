package fixture

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"testing"
	"time"
)

func DrainStreamEvents(
	t *testing.T,
	events <-chan *stream_event.Event,
) {
	t.Helper()
	deadline := time.After(5 * time.Second)

	for {
		select {
		case _, open := <-events:
			if !open {
				return
			}
		case <-deadline:
			t.Fatal("timed out waiting for the stream to close")

			return
		}
	}
}
