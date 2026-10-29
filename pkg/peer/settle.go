package peer

import (
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (g *Guard) settle(
	trial bool,
	e error,
) (*edge.Edge, error) {
	crossed, refused := g.reacher.Triage(g.name, e)
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if trial {
		g.trying = false
	}

	switch {
	case refused != nil:
		if trial {
			crossed = g.reacher.Succeed(g.name)
			g.served = false
		}
	case unavailable.Is(e):
		if crossed != nil {
			g.served = true
			g.delay = g.minimum
			g.retry = g.clock().Add(g.delay)
		} else if trial {
			g.delay = min(2*g.delay, g.maximum)
			g.retry = g.clock().Add(g.delay)
		}
	default:
		g.served = false
	}

	return crossed, refused
}
