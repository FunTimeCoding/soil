package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"

func (c *Client) MergeRequest(
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	result, _, e := c.client.MergeRequests.GetMergeRequest(
		project,
		identifier,
		nil,
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return merge_request_detail.New(result), nil
}
