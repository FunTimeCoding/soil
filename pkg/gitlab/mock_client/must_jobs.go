package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) MustJobs(x context.Context) []*job.Job {
	result, e := c.Jobs(x)
	errors.PanicOnError(e)

	return result
}
