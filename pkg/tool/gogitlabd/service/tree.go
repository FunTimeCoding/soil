package service

import "github.com/funtimecoding/soil/pkg/gitlab/tree"

func (s *Service) Tree(
	instance string,
	project int64,
	path string,
	reference string,
	recursive bool,
	limit int64,
) ([]*tree.Node, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.Tree(project, path, reference, recursive, limit)
}
