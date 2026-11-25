package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) observe(
	x context.Context,
	k *watched_request.Key,
) (*snapshot.Snapshot, *merge_request_detail.Detail, error) {
	d, e := s.MergeRequest(x, k.Instance, k.Project, k.Request)

	if e != nil {
		return nil, nil, e
	}

	approvers, f := guarded(
		x,
		s,
		k.Instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) ([]string, error) {
			return c.MergeRequestApprovals(y, k.Project, k.Request)
		},
	)

	if f != nil {
		return nil, nil, f
	}

	return snapshot.New(k, d, approvers), d, nil
}
