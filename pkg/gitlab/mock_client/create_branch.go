package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
)

func (c *Client) CreateBranch(
	_ context.Context,
	_ int64,
	_ string,
	_ string,
) (*branch.Branch, error) {
	return nil, nil
}
