package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"github.com/funtimecoding/soil/pkg/gitlab/types/recorded_commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Commit(
	_ int64,
	branch string,
	message string,
	_ string,
	_ string,
	_ bool,
) (*commit.Commit, error) {
	c.commits = append(c.commits, recorded_commit.New(branch, message))

	return commit.New(&gitlab.Commit{}), nil
}
