package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustRunnerJobs(
	x context.Context,
	runner int64,
	stopAfter int,
) []*job.Job {
	result, e := c.RunnerJobs(x, runner, stopAfter)
	errors.PanicOnError(e)

	return result
}
