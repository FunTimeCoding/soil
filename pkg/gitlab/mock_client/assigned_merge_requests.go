package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) AssignedMergeRequests(
	_ context.Context,
	_ bool,
) ([]*merge_request.Request, error) {
	return c.assignedRequests, nil
}
