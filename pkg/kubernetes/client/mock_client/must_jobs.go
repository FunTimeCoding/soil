package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"

func (c *Client) MustJobs(_ string) []*job.Job {
	return nil
}
