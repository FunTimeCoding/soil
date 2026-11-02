package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) ProjectsMergeRequests(
	_ context.Context,
) ([]*merge_request.Request, error) {
	return nil, c.failure
}
