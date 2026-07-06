package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) enrichProjectJob(
	j *job.Job,
	p *project.Project,
) *job.Job {
	j.Project = p
	c.enrichJobCommon(j)

	return j
}
