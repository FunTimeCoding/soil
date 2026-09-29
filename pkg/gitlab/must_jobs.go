package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustJobs() []*job.Job {
	result, e := c.Jobs()
	errors.PanicOnError(e)

	return result
}
