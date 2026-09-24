package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) AssignedMergeRequests(_ bool) (
	[]*merge_request.Request,
	error,
) {
	return c.assignedRequests, nil
}
