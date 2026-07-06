package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) MustRunners(all bool) []*runner.Runner {
	result, e := c.Runners(all)
	errors.PanicOnError(e)

	return result
}
