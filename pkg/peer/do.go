package peer

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
)

func (g *Guard) Do(
	x context.Context,
	call func(context.Context) error,
) error {
	e := g.Attempt(x, call)

	if e == nil || x.Err() != nil || unavailable.Is(e) {
		return e
	}

	g.handle(e)

	return e
}
