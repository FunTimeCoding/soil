package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) AssignedReviews(
	_ context.Context,
	_ bool,
) ([]*merge_request.Request, error) {
	return nil, c.failure
}
