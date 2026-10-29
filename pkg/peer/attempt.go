package peer

import "context"

func (g *Guard) Attempt(
	x context.Context,
	call func(context.Context) error,
) error {
	trial, e := g.admit()

	if e != nil {
		return e
	}

	f := call(x)

	if x.Err() != nil {
		g.abandon(trial)

		return f
	}

	crossed, refused := g.settle(trial, f)

	if crossed != nil {
		g.logger.Plain("%s", crossed)
	}

	if h := g.Unavailable(); h != nil {
		g.awaken()

		return h
	}

	return refused
}
