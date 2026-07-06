package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(p *project.Project) ([]*job.Job, error) {
	result, _, e := c.client.Jobs.ListProjectJobs(p.Identifier, nil)

	if e != nil {
		return nil, e
	}

	return c.enrichProjectJobs(job.NewSlice(result), p), nil
}
