package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustRetryJob(
	x context.Context,
	j *job.Job,
) *job.Job {
	result, e := c.RetryJob(x, j)
	errors.PanicOnError(e)

	return result
}
