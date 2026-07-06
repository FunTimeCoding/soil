package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustRetry(
	project int64,
	jobIdentifier int64,
) *job.Job {
	result, e := c.Retry(project, jobIdentifier)
	errors.PanicOnError(e)

	return result
}
