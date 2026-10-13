package service

import "github.com/funtimecoding/soil/pkg/gitlab/file"

func (s *Service) File(
	instance string,
	project int64,
	branch string,
	name string,
) (*file.File, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.File(project, branch, name)
}
