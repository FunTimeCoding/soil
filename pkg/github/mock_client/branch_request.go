package mock_client

import "github.com/funtimecoding/soil/pkg/github/pull_request"

func (c *Client) BranchRequest(_ string, _ string, _ string) (*pull_request.Request, error) {
	return nil, c.Failure
}
