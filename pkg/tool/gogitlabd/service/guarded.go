package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func guarded[T any](
	x context.Context,
	s *Service,
	instance string,
	f func(context.Context, face.GitLabClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	g := p.Guard().Attempt(
		x,
		func(y context.Context) error {
			var h error
			result, h = f(y, p.Client())

			return h
		},
	)

	return result, g
}
