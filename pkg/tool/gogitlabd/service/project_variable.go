package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (s *Service) ProjectVariable(
	x context.Context,
	instance string,
	project int64,
	key string,
) (*variable.Variable, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*variable.Variable, error) {
			return c.ProjectVariable(y, project, key)
		},
	)
}
