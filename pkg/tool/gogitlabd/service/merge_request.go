package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func (s *Service) MergeRequest(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*merge_request_detail.Detail, error) {
			return c.MergeRequest(y, project, identifier)
		},
	)
}
