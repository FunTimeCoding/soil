package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) ReviewingMergeRequests(_ bool) (
	[]*merge_request.Request,
	error,
) {
	return c.reviewingRequests, nil
}
