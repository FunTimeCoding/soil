package mock_client

import "github.com/funtimecoding/soil/pkg/github/pull_request"

func (c *Client) FindBranchRequest(_ string, _ string, _ string) (*pull_request.Request, bool, error) {
	return nil, false, c.Failure
}
