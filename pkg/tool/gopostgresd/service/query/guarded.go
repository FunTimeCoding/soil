package query

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

func (s *Service) guarded(
	x context.Context,
	instance string,
	sql string,
) (*result.Result, error) {
	p, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return p.Guarded(x, theirs, sql)
}
