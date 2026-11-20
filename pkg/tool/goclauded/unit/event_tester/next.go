package event_tester

import "time"

func (o *Tester) Next(events <-chan Event) *Event {
	o.t.Helper()

	select {
	case e := <-events:
		return &e
	case <-time.After(5 * time.Second):
		o.t.Fatal("timed out waiting for SSE event")

		return nil
	}
}
