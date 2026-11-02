package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (s *Service) Projects(
	instance string,
	search string,
) ([]*project.Project, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*project.Project, error) {
			if search != "" {
				return c.SearchProject(search)
			}

			return c.Projects()
		},
	)
}
