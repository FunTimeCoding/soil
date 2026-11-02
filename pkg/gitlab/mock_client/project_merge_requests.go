package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) ProjectMergeRequests(
	_ context.Context,
	_ int64,
	_ string,
) ([]*merge_request.Request, error) {
	return nil, nil
}
