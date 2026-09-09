package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (c *Client) ListCommits(
	_ int64,
	_ string,
	_ int64,
) ([]*commit.Commit, error) {
	return nil, nil
}
