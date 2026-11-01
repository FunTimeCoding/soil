package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/commit"

func (c *Client) MergeRequestCommits(_ int64, _ int64) ([]*commit.Commit, error) {
	return nil, c.failure
}
