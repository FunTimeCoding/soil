package peer

import "github.com/funtimecoding/soil/pkg/errors/unavailable"

func (g *Guard) Unavailable() error {
	down, h := g.reacher.Down(g.name)

	if !down {
		return nil
	}

	return unavailable.New(g.name, h.Since, h.Reason)
}
