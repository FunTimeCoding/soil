package peer

import "context"

type Guard struct{}

func (g *Guard) Attempt(
	x context.Context,
	call func(context.Context) error,
) error {
	return call(x)
}

func (g *Guard) Do(
	x context.Context,
	call func(context.Context) error,
) error {
	return call(x)
}

func (g *Guard) Name() string {
	return "catalog"
}
