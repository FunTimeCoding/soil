package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
)

func (c *Client) Branches(
	_ context.Context,
	_ int64,
) ([]*branch.Branch, error) {
	return c.branches, nil
}
