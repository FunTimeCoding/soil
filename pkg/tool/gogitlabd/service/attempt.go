package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) attempt(
	instance string,
	f func(face.GitLabClient) error,
) error {
	p, e := s.pool.Client(instance)

	if e != nil {
		return e
	}

	return p.Guard().Attempt(
		context.Background(),
		func(context.Context) error {
			return f(p.Client())
		},
	)
}
