package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) SetMergeRequests(
	assigned []*merge_request.Request,
	reviewing []*merge_request.Request,
) {
	c.assignedRequests = assigned
	c.reviewingRequests = reviewing
}
