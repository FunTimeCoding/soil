package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
)

func (c *Client) MergeRequest(
	_ context.Context,
	_ int64,
	_ int64,
) (*merge_request_detail.Detail, error) {
	return nil, nil
}
