package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) MustRunnersByTag(tag string) []*runner.Runner {
	result, e := c.RunnersByTag(tag)
	errors.PanicOnError(e)

	return result
}
