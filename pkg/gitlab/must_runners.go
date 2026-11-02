package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) MustRunners(
	x context.Context,
	all bool,
) []*runner.Runner {
	result, e := c.Runners(x, all)
	errors.PanicOnError(e)

	return result
}
