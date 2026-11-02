package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (s *Service) Projects(
	x context.Context,
	instance string,
	search string,
) ([]*project.Project, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*project.Project, error) {
			if search != "" {
				return c.SearchProject(y, search)
			}

			return c.Projects(y)
		},
	)
}
