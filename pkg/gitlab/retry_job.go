package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RetryJob(
	x context.Context,
	j *job.Job,
) (*job.Job, error) {
	if j.Project == nil {
		return nil, unexpected.Format("job has no project")
	}

	return c.Retry(x, j.Project.Identifier, j.Identifier)
}
