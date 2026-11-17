package peer

import "time"

func (g *Guard) WithClock(clock func() time.Time) *Guard {
	g.clock = clock
	g.outages.WithClock(clock)

	return g
}
