package service

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (s *Service) Branches(
	instance string,
	project int64,
) ([]*branch.Branch, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Branches(project)
}
