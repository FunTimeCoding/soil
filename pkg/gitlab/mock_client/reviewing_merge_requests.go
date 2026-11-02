package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) ReviewingMergeRequests(
	_ context.Context,
	_ bool,
) ([]*merge_request.Request, error) {
	return c.reviewingRequests, nil
}
