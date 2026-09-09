package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CommitActions(
	_ int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) (*commit.Commit, error) {
	c.commits = append(
		c.commits,
		&RecordedCommit{Branch: branch, Message: message, Actions: v},
	)

	return commit.New(&gitlab.Commit{}), nil
}
