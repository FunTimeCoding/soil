package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) MergeRequestDiscussions(
	project int64,
	mergeRequest int64,
) ([]*gitlab.Discussion, error) {
	result, _, e := c.client.Discussions.ListMergeRequestDiscussions(
		project,
		mergeRequest,
		nil,
	)

	return result, wrapError(e)
}
