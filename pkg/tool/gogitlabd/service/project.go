package service

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (s *Service) Project(
	instance string,
	identifier int64,
) (*project.Project, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Project(identifier)
}
