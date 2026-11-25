package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"slices"
)

func (s *Service) MergeRequestsOf(
	x context.Context,
	instance string,
	f *merge_request_filter.Filter,
) ([]*merge_request.Request, error) {
	switch {
	case f.Limit < 1 || f.Limit > gitlabConstant.RequestLimitMaximum:
		return nil, validation.New(
			"limit must be between 1 and %d: %d",
			gitlabConstant.RequestLimitMaximum,
			f.Limit,
		)
	case f.Offset < 0:
		return nil, validation.New("offset must not be negative: %d", f.Offset)
	case f.State != gitlabConstant.AllState &&
		!slices.Contains(gitlabConstant.RequestStates, f.State):

		return nil, validation.New("unknown state: %s", f.State)
	case f.OrderBy != gitlabConstant.UpdatedOrder && f.OrderBy != gitlabConstant.CreatedOrder:
		return nil, validation.New("unknown order: %s", f.OrderBy)
	}

	return guarded(
		x,
		s,
		instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) ([]*merge_request.Request, error) {
			return c.MergeRequests(y, f)
		},
	)
}
