package peer

import "github.com/funtimecoding/soil/pkg/peer/outage_tracker/edge"

func (g *Guard) handle(e error) {
	g.outages.Handle(
		g.name,
		e,
		func(x *edge.Edge) { g.logger.Plain("%s", x) },
		func(f error) { g.reporter.CaptureException(f) },
	)

	if g.Unavailable() != nil {
		g.awaken()
	}
}
