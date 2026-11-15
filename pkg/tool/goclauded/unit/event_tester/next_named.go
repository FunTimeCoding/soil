package event_tester

import "testing"

func NextNamed(
	t *testing.T,
	events <-chan Event,
	name string,
) *Event {
	t.Helper()

	for {
		if e := Next(t, events); e.Name == name {
			return e
		}
	}
}
