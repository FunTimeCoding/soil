package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (s *Service) ProjectMergeRequests(
	instance string,
	project int64,
	state string,
) ([]*merge_request.Request, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*merge_request.Request, error) {
			return c.ProjectMergeRequests(project, state)
		},
	)
}
