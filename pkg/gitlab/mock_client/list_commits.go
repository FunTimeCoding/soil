package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) ListCommits(
	_ context.Context,
	_ int64,
	_ string,
	_ int64,
) ([]*commit.Commit, error) {
	return nil, nil
}
