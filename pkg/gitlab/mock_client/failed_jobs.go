package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) FailedJobs(_ context.Context) ([]*job.Job, error) {
	return nil, c.failure
}
