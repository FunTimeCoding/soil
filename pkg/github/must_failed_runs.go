package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (c *Client) MustFailedRuns(verbose bool) []*run.Run {
	result, e := c.FailedRuns(verbose)
	errors.PanicOnError(e)

	return result
}
