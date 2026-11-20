package event_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stream_event"
	"time"
)

func (o *Tester) Drain(events <-chan *stream_event.Event) {
	o.t.Helper()
	deadline := time.After(5 * time.Second)

	for {
		select {
		case _, open := <-events:
			if !open {
				return
			}
		case <-deadline:
			o.t.Fatal("timed out waiting for the stream to close")

			return
		}
	}
}
