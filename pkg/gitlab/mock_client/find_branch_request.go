package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) FindBranchRequest(_ int64, _ string) (*merge_request.Request, bool, error) {
	return nil, false, c.failure
}
