package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) PipelineJobs(
	_ context.Context,
	_ int64,
	_ int64,
) ([]*job.Job, error) {
	return c.jobs, nil
}
