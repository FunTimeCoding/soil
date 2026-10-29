package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"

func (c *Client) CronJob(_ string, _ string) (*cron_job.Job, error) {
	return nil, c.Failure
}
