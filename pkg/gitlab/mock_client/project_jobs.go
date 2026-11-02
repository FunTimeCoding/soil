package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(
	_ context.Context,
	_ *project.Project,
) ([]*job.Job, error) {
	return nil, c.failure
}
