package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (s *Service) Project(
	x context.Context,
	instance string,
	identifier int64,
) (*project.Project, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*project.Project, error) {
			return c.Project(y, identifier)
		},
	)
}
