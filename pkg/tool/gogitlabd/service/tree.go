package service

import "github.com/funtimecoding/soil/pkg/gitlab/tree"

func (s *Service) Tree(
	project int64,
	path string,
	reference string,
	recursive bool,
	limit int64,
) ([]*tree.Node, error) {
	return s.client.Tree(project, path, reference, recursive, limit)
}
