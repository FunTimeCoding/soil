package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) DeletePipeline(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) error {
	return s.attempt(
		x,
		instance,
		func(y context.Context, c face.GitLabClient) error {
			return c.DeletePipeline(y, project, identifier)
		},
	)
}
