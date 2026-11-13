package peer

import "context"

func Guarded[T any](
	x context.Context,
	g *Guard,
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
