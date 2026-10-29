package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"
)

func (c *Client) CronJobs(namespace string) ([]*cron_job.Job, error) {
	result, e := list.CronJob(c.client, c.context, namespace, "")

	if e != nil {
		return nil, e
	}

	return cron_job.NewSlice(result, c.cluster), nil
}
