package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
)

func (c *Client) MustBranchRequest(
	x context.Context,
	project int64,
	branch string,
) *merge_request.Request {
	result, e := c.BranchRequest(x, project, branch)
	errors.PanicOnError(e)

	return result
}
