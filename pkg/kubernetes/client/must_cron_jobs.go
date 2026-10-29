package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"
)

func (c *Client) MustCronJobs(namespace string) []*cron_job.Job {
	result, e := c.CronJobs(namespace)
	errors.PanicOnError(e)

	return result
}
