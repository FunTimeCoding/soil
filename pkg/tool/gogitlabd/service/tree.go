package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
)

func (s *Service) Tree(
	x context.Context,
	instance string,
	project int64,
	path string,
	reference string,
	recursive bool,
	limit int64,
) ([]*tree.Node, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*tree.Node, error) {
			return c.Tree(y, project, path, reference, recursive, limit)
		},
	)
}
