package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RetryJob(j *job.Job) (*job.Job, error) {
	if j.Project == nil {
		return nil, unexpected.Format("job has no project")
	}

	return c.Retry(j.Project.Identifier, j.Identifier)
}
