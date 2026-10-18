package peer

import "context"

func (g *Guard) Wait(x context.Context) error {
	e := Await(x, g.logger, g.name, g.address, g.probe.Reachable)

	if x.Err() == nil {
		g.handle(e)
	}

	return e
}
