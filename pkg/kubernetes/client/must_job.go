package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"
)

func (c *Client) MustJob(
	namespace string,
	name string,
) *job.Job {
	result, e := c.Job(namespace, name)
	errors.PanicOnError(e)

	return result
}
