package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RetryJob(
	_ context.Context,
	_ *job.Job,
) (*job.Job, error) {
	return nil, c.failure
}
