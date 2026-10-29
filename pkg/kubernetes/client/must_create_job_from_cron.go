package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"k8s.io/api/batch/v1"
)

func (c *Client) MustCreateJobFromCron(
	namespace string,
	cron string,
	name string,
) *v1.Job {
	result, e := c.CreateJobFromCron(namespace, cron, name)
	errors.PanicOnError(e)

	return result
}
