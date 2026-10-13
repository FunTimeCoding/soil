package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) ProjectVariable(
	instance string,
	project int64,
	key string,
) (*variable.Variable, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.ProjectVariable(project, key)
}
