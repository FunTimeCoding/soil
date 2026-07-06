package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) RetryJob(j *job.Job) (*job.Job, error) {
	if j.Project == nil {
		return nil, fmt.Errorf("job has no project")
	}

	return c.Retry(j.Project.Identifier, j.Identifier)
}
