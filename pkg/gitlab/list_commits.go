package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) ListCommits(
	project int64,
	o *gitlab.ListCommitsOptions,
) ([]*gitlab.Commit, error) {
	result, _, e := c.client.Commits.ListCommits(project, o)

	return result, wrapError(e)
}
