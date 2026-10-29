package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/cron_job"

func (c *Client) MustCronJobs(_ string) []*cron_job.Job {
	return nil
}
