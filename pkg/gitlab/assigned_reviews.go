package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) AssignedReviews(
	x context.Context,
	all bool,
) ([]*merge_request.Request, error) {
	u, e := c.userIdentifier(x)

	if e != nil {
		return nil, e
	}

	requests, _, f := c.client.MergeRequests.ListMergeRequests(
		&gitlab.ListMergeRequestsOptions{
			ReviewerID:  gitlab.ReviewerID(u),
			State:       new(constant.OpenedState),
			ListOptions: constant.DefaultListOptions,
		},
		gitlab.WithContext(x),
	)

	if f != nil {
		return nil, wrapError(f)
	}

	result := merge_request.NewSlice(requests)

	if all {
		return result, nil
	}

	return merge_request.FilterDone(result), nil
}
