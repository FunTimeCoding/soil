package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) AssignedMergeRequests(all bool) ([]*merge_request.Request, error) {
	u, e := c.userIdentifier()

	if e != nil {
		return nil, e
	}

	requests, _, f := c.client.MergeRequests.ListMergeRequests(
		&gitlab.ListMergeRequestsOptions{
			AssigneeID:  gitlab.AssigneeID(u),
			State:       new(constant.OpenedState),
			ListOptions: constant.DefaultListOptions,
		},
		nil,
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
