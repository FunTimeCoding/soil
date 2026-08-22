package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) ReadCommit(
	project int64,
	sha string,
) (*gitlab.Commit, error) {
	result, _, e := c.client.Commits.GetCommit(
		project,
		sha,
		&gitlab.GetCommitOptions{},
	)

	return result, wrapError(e)
}
