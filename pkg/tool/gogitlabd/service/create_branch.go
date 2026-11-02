package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) CreateBranch(
	x context.Context,
	instance string,
	project int64,
	name string,
	reference string,
) (*branch.Branch, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*branch.Branch, error) {
			return c.CreateBranch(y, project, name, reference)
		},
	)
}
