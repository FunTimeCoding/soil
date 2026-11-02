package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) DeleteProjectVariable(
	x context.Context,
	instance string,
	project int64,
	key string,
) error {
	return s.attempt(
		x,
		instance,
		func(y context.Context, c face.GitLabClient) error {
			return c.DeleteProjectVariable(y, project, key)
		},
	)
}
