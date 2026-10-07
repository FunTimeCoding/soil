package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/types/recorded_commit"

func (c *Client) Commits() []*recorded_commit.Commit {
	return c.commits
}
