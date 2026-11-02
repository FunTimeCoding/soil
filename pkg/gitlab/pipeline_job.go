package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) PipelineJob(
	x context.Context,
	project int64,
	identifier int64,
) (*job.Job, error) {
	result, _, e := c.client.Jobs.GetJob(
		project,
		identifier,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichJob(x, job.New(result))
}
