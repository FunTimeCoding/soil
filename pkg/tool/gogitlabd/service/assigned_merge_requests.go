package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (s *Service) AssignedMergeRequests(
	instance string,
	all bool,
) ([]*merge_request.Request, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) ([]*merge_request.Request, error) {
			return c.AssignedMergeRequests(all)
		},
	)
}
