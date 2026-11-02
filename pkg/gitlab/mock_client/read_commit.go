package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) ReadCommit(
	_ context.Context,
	_ int64,
	_ string,
) (*commit.Commit, error) {
	return nil, nil
}
