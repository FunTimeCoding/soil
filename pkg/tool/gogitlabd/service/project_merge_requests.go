package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (s *Service) ProjectMergeRequests(
	x context.Context,
	instance string,
	project int64,
	state string,
) ([]*merge_request.Request, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) ([]*merge_request.Request, error) {
			return c.ProjectMergeRequests(y, project, state)
		},
	)
}
