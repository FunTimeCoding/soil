package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Retry(
	x context.Context,
	project int64,
	jobIdentifier int64,
) (*job.Job, error) {
	result, _, e := c.client.Jobs.RetryJob(
		project,
		jobIdentifier,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichJob(x, job.New(result))
}
