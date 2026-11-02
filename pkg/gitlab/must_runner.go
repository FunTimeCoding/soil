package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) MustRunner(
	x context.Context,
	identifier int64,
) *runner.Runner {
	result, e := c.Runner(x, identifier)
	errors.PanicOnError(e)

	return result
}
