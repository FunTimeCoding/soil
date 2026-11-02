package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
)

func (c *Client) MergeRequestCommits(
	_ context.Context,
	_ int64,
	_ int64,
) ([]*commit.Commit, error) {
	return nil, c.failure
}
