package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) MergeRequestDiffs(
	project int64,
	mergeRequest int64,
) ([]*gitlab.MergeRequestDiff, error) {
	result, _, e := c.client.MergeRequests.ListMergeRequestDiffs(
		project,
		mergeRequest,
		nil,
	)

	return result, wrapError(e)
}
