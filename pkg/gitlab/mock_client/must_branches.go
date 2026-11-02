package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
)

func (c *Client) MustBranches(
	_ context.Context,
	_ int64,
) []*branch.Branch {
	return c.branches
}
