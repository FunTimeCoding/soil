package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustFailedJobs() []*job.Job {
	result, e := c.FailedJobs()
	errors.PanicOnError(e)

	return result
}
