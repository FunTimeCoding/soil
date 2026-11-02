package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustProjectsMergeRequests(
	x context.Context,
) []*merge_request.Request {
	result, e := c.ProjectsMergeRequests(x)
	errors.PanicOnError(e)

	return result
}
