package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/commit"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CommitActions(
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
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return commit.New(result), nil
}
