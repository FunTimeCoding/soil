package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
)

func (s *Service) Trace(
	x context.Context,
	instance string,
	project int64,
	jobIdentifier int64,
) (string, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (string, error) {
			return c.Trace(y, project, jobIdentifier)
		},
	)
}
