package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) Variables(
	instance string,
	project int64,
) ([]*variable.Variable, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Variables(project)
}
