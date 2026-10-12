package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) ProjectVariable(
	project int64,
	key string,
) (*variable.Variable, error) {
	return s.client.ProjectVariable(project, key)
}
