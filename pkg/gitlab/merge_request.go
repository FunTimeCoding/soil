package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequest(
	x context.Context,
	project int64,
	identifier int64,
) (*merge_request_detail.Detail, error) {
	result, _, e := c.client.MergeRequests.GetMergeRequest(
		project,
		identifier,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return merge_request_detail.New(result), nil
}
