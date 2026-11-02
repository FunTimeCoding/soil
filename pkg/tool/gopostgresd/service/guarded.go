package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

func (s *Service) guarded(
	x context.Context,
	instance string,
	narrow func(context.Context, error) error,
	sql string,
	a ...any,
) (*result.Result, error) {
	p, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	var v *result.Result
	f := p.Guard().Attempt(
		x,
		func(y context.Context) error {
			var g error
			v, g = p.Client().Query(y, sql, a...)

			return g
		},
	)

	if f != nil {
		return nil, narrow(x, f)
	}

	return v, nil
}
