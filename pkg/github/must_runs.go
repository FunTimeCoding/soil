package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (c *Client) MustRuns(
	loadJobs bool,
	verbose bool,
) []*run.Run {
	result, e := c.Runs(loadJobs, verbose)
	errors.PanicOnError(e)

	return result
}
