package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
)

func (c *Client) MustJobs(namespace string) []*job.Job {
	result, e := c.Jobs(namespace)
	errors.PanicOnError(e)

	return result
}
