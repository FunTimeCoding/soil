package peered

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

func (p *Peered) Guarded(
	x context.Context,
	narrow func(context.Context, error) error,
	sql string,
	a ...any,
) (*result.Result, error) {
	var v *result.Result
	f := p.guard.Attempt(
		x,
		func(y context.Context) error {
			var g error
			v, g = p.client.Query(y, sql, a...)

			return g
		},
	)

	if f != nil {
		return nil, narrow(x, f)
	}

	return v, nil
}
