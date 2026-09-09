package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Commit(
	_ int64,
	branch string,
	message string,
	_ string,
	_ string,
	_ bool,
) (*commit.Commit, error) {
	c.commits = append(
		c.commits,
		&RecordedCommit{Branch: branch, Message: message},
	)

	return commit.New(&gitlab.Commit{}), nil
}
