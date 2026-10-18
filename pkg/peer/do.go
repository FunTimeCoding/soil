package peer

import "context"

func (g *Guard) Do(
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

	g.handle(e)

	if f := g.Unavailable(); f != nil {
		return f
	}

	return e
}
