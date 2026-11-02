package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RunnerJobs(
	_ context.Context,
	_ int64,
	_ int,
) ([]*job.Job, error) {
	return nil, c.failure
}
