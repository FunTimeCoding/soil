package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func guarded[T any](
	g *peer.Guard,
	call func() (T, error),
) (T, error) {
	var result T
	e := g.Attempt(
		context.Background(),
		func(context.Context) error {
			var f error
			result, f = call()

			return f
		},
	)

	return result, e
}
