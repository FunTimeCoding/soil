package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) MergeRequest(
	project int64,
	mergeRequest int64,
) (*gitlab.MergeRequest, error) {
	result, _, e := c.client.MergeRequests.GetMergeRequest(
		project,
		mergeRequest,
		nil,
	)

	return result, wrapError(e)
}
