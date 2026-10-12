package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (s *Service) ProjectMergeRequests(
	project int64,
	state string,
) ([]*merge_request.Request, error) {
	return s.client.ProjectMergeRequests(project, state)
}
