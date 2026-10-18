package peer

import "time"

func (g *Guard) WithBackoff(
	minimum time.Duration,
	maximum time.Duration,
) *Guard {
	g.minimum = minimum
	g.maximum = maximum

	return g
}
