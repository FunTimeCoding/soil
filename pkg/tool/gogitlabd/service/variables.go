package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/variable"
)

func (s *Service) Variables(
	instance string,
	project int64,
) ([]*variable.Variable, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*variable.Variable, error) {
			return c.Variables(project)
		},
	)
}
