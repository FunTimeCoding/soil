package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CommitDiff(
	project int64,
	sha string,
) ([]*gitlab.Diff, error) {
	result, _, e := c.client.Commits.GetCommitDiff(project, sha, nil)

	return result, wrapError(e)
}
