package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustRetryJob(j *job.Job) *job.Job {
	result, e := c.RetryJob(j)
	errors.PanicOnError(e)

	return result
}
