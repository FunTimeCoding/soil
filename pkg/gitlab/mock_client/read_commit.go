package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (c *Client) ReadCommit(
	_ int64,
	_ string,
) (*commit.Commit, error) {
	return nil, nil
}
