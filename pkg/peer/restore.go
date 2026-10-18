package peer

import "time"

func (g *Guard) restore() {
	delay := g.minimum

	for g.Unavailable() != nil {
		select {
		case <-g.context.Done():
			return
		case <-time.After(delay):
		}

		e := g.probe.Reachable(g.context)

		if g.context.Err() != nil {
			return
		}

		g.handle(e)
		delay = min(2*delay, g.maximum)
	}
}
