package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

func (s *Service) Query(
	x context.Context,
	instance string,
	sql string,
) (*result.Result, error) {
	return s.guarded(x, instance, theirs, sql)
}
