package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
)

func (s *Service) Tree(
	instance string,
	project int64,
	path string,
	reference string,
	recursive bool,
	limit int64,
) ([]*tree.Node, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*tree.Node, error) {
			return c.Tree(project, path, reference, recursive, limit)
		},
	)
}
