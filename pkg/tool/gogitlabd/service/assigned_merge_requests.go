package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (s *Service) AssignedMergeRequests(
	all bool,
) ([]*merge_request.Request, error) {
	return s.client.AssignedMergeRequests(all)
}
