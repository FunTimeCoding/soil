package fixture

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"testing"
	"time"
)

func NextStreamEvent(
	t *testing.T,
	events <-chan *stream_event.Event,
) *stream_event.Event {
	t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for stream event")

		return nil
	}
}
