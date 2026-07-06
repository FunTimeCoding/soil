package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (c *Client) MustProjectRuns(
	owner string,
	name string,
) []*run.Run {
	result, e := c.ProjectRuns(owner, name)
	errors.PanicOnError(e)

	return result
}
