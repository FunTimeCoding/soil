package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(_ *project.Project) ([]*job.Job, error) {
	return nil, c.failure
}
