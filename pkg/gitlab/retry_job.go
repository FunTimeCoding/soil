package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RetryJob(j *job.Job) (*job.Job, error) {
	if j.Project == nil {
		return nil, validation.New("job has no project")
	}

	return c.Retry(j.Project.Identifier, j.Identifier)
}
