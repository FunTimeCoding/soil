package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ReadCommit(
	x context.Context,
	project int64,
	sha string,
) (*commit.Commit, error) {
	result, _, e := c.client.Commits.GetCommit(
		project,
		sha,
		&gitlab.GetCommitOptions{},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.New(result), nil
}
