package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/types/recorded_commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CommitActions(
	_ int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) (*commit.Commit, error) {
	o := recorded_commit.New(branch, message)
	o.Actions = v
	c.commits = append(c.commits, o)

	return commit.New(&gitlab.Commit{}), nil
}
