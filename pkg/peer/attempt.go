package peer

import "context"

func (g *Guard) Attempt(
	x context.Context,
	call func(context.Context) error,
) error {
	if e := g.Unavailable(); e != nil {
		return e
	}

	e := call(x)

	if x.Err() != nil {
		return e
	}

	crossed, refused := g.reacher.Triage(g.name, e)

	if crossed != nil {
		g.logger.Plain("%s", crossed)
	}

	if f := g.Unavailable(); f != nil {
		g.awaken()

		return f
	}

	return refused
}
