package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) ReadCommit(
	project int64,
	sha string,
) (*commit.Commit, error) {
	result, _, e := c.client.Commits.GetCommit(
		project,
		sha,
		&gitlab.GetCommitOptions{},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.New(result), nil
}
