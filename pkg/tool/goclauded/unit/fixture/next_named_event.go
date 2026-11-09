package fixture

import "testing"

func NextNamedEvent(
	t *testing.T,
	events <-chan Event,
	name string,
) *Event {
	t.Helper()

	for {
		if e := NextEvent(t, events); e.Name == name {
			return e
		}
	}
}
