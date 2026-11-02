package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (s *Service) ProjectVariable(
	instance string,
	project int64,
	key string,
) (*variable.Variable, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*variable.Variable, error) {
			return c.ProjectVariable(project, key)
		},
	)
}
