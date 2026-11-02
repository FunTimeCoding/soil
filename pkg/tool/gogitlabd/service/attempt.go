package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) attempt(
	x context.Context,
	instance string,
	f func(context.Context, face.GitLabClient) error,
) error {
	p, e := s.pool.Client(instance)

	if e != nil {
		return e
	}

	return p.Guard().Attempt(
		x,
		func(y context.Context) error {
			return f(y, p.Client())
		},
	)
}
