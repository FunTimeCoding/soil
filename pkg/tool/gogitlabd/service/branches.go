package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) Branches(
	x context.Context,
	instance string,
	project int64,
) ([]*branch.Branch, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*branch.Branch, error) {
			return c.Branches(y, project)
		},
	)
}
