package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
)

func (c *Client) MustGraphRunner(
	x context.Context,
	identifier int64,
) *response.Runner {
	result, e := c.GraphRunner(x, identifier)
	errors.PanicOnError(e)

	return result
}
