package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func (s *Service) MergeRequest(
	instance string,
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*merge_request_detail.Detail, error) {
			return c.MergeRequest(project, identifier)
		},
	)
}
