package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CommitActions(
	x context.Context,
	project int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) (*commit.Commit, error) {
	result, _, e := c.client.Commits.CreateCommit(
		project,
		&gitlab.CreateCommitOptions{
			Branch:        &branch,
			CommitMessage: &message,
			Actions:       v,
		},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.New(result), nil
}
