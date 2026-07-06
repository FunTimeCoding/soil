package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) MustRunner(identifier int64) *runner.Runner {
	result, e := c.Runner(identifier)
	errors.PanicOnError(e)

	return result
}
