package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CommitActions(
	project int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) (*gitlab.Commit, error) {
	result, _, e := c.client.Commits.CreateCommit(
		project,
		&gitlab.CreateCommitOptions{
			Branch:        &branch,
			CommitMessage: &message,
			Actions:       v,
		},
	)

	return result, wrapError(e)
}
