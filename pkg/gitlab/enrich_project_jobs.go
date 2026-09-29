package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) enrichProjectJobs(
	v []*job.Job,
	p *project.Project,
) ([]*job.Job, error) {
	for _, j := range v {
		if _, e := c.enrichProjectJob(j, p); e != nil {
			return nil, e
		}
	}

	return v, nil
}
