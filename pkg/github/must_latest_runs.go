package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/run"
)

func (c *Client) MustLatestRuns(owner, repo string) []*run.Run {
	result, e := c.LatestRuns(owner, repo)
	errors.PanicOnError(e)

	return result
}
