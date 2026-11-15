package guard_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/peer"
)

func Panicked(
	x context.Context,
	g *peer.Guard,
) (result any) {
	defer func() { result = recover() }()
	errors.PanicOnError(
		g.Attempt(x, func(context.Context) error { panic("trial") }),
	)

	return nil
}
