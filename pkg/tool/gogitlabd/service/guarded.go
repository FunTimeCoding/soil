package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func guarded[T any](
	s *Service,
	instance string,
	f func(face.GitLabClient) (T, error),
) (T, error) {
	var result T
	p, e := s.pool.Client(instance)

	if e != nil {
		return result, e
	}

	g := p.Guard().Attempt(
		context.Background(),
		func(context.Context) error {
			var h error
			result, h = f(p.Client())

			return h
		},
	)

	return result, g
}
