package service

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (s *Service) CreateBranch(
	instance string,
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.CreateBranch(project, name, reference)
}
