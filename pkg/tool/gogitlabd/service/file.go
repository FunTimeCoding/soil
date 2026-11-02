package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
)

func (s *Service) File(
	x context.Context,
	instance string,
	project int64,
	branch string,
	name string,
) (*file.File, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*file.File, error) {
			return c.File(y, project, branch, name)
		},
	)
}
