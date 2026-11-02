package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) ResolveProject(
	x context.Context,
	instance string,
	identifier string,
) (int64, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (int64, error) {
			return c.ResolveProject(y, identifier)
		},
	)
}
