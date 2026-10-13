package service

import "github.com/funtimecoding/soil/pkg/gitlab/tag"

func (s *Service) Tags(
	instance string,
	project int64,
) ([]*tag.Tag, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Tags(project)
}
