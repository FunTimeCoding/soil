package event_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"time"
)

func (o *Tester) NextStream(
	events <-chan *stream_event.Event,
) *stream_event.Event {
	o.t.Helper()

	select {
	case e := <-events:
		return e
	case <-time.After(5 * time.Second):
		o.t.Fatal("timed out waiting for stream event")

		return nil
	}
}
