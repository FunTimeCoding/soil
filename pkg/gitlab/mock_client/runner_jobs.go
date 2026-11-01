package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) RunnerJobs(_ int64, _ int) ([]*job.Job, error) {
	return nil, c.failure
}
