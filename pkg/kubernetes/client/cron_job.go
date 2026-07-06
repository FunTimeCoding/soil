package client

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/client/operation/get"
	"github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"
)

func (c *Client) CronJob(
	namespace string,
	name string,
) *cron_job.Job {
	result := get.CronJob(c.client, c.context, namespace, name)

	if result == nil {
		return nil
	}

	return cron_job.New(result, c.cluster)
}
