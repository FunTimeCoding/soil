package service

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (s *Service) Projects(
	instance string,
	search string,
) ([]*project.Project, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	if search != "" {
		return c.SearchProject(search)
	}

	return c.Projects()
}
