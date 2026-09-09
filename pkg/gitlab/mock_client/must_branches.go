package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/branch"

func (c *Client) MustBranches(_ int64) []*branch.Branch {
	return nil
}
