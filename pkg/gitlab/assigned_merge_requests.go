package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) AssignedMergeRequests(all bool) ([]*merge_request.Request, error) {
	requests, _, e := c.client.MergeRequests.ListMergeRequests(
		&gitlab.ListMergeRequestsOptions{
			AssigneeID:  gitlab.AssigneeID(c.user.ID),
			State:       new(constant.OpenedState),
			ListOptions: constant.DefaultListOptions,
		},
		nil,
	)

	if e != nil {
		return nil, e
	}

	result := merge_request.NewSlice(requests)

	if all {
		return result, nil
	}

	return merge_request.FilterDone(result), nil
}
