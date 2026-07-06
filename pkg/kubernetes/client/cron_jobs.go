package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"
)

func (c *Client) CronJobs(namespace string) []*cron_job.Job {
	return cron_job.NewSlice(
		list.CronJob(c.client, c.context, namespace, ""),
		c.cluster,
	)
}
