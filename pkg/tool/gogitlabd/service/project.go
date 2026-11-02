package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (s *Service) Project(
	instance string,
	identifier int64,
) (*project.Project, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*project.Project, error) {
			return c.Project(identifier)
		},
	)
}
