package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) ProjectsMergeRequests() ([]*merge_request.Request, error) {
	return nil, c.failure
}
