package service

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"

func (s *Service) MergeRequest(
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	return s.client.MergeRequest(project, identifier)
}
