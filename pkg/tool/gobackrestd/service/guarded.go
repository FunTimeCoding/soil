package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func guarded[T any](
	x context.Context,
	g *peer.Guard,
	call func(context.Context) (T, error),
) (T, error) {
	var result T
	e := g.Attempt(
		x,
		func(y context.Context) error {
			var f error
			result, f = call(y)

			return f
		},
	)

	return result, e
}
