package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) enrichProjectJobs(
	v []*job.Job,
	p *project.Project,
) []*job.Job {
	for _, j := range v {
		c.enrichProjectJob(j, p)
	}

	return v
}
