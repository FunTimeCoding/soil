package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) enrichProjectJob(
	j *job.Job,
	p *project.Project,
) (*job.Job, error) {
	j.Project = p

	if e := c.enrichJobCommon(j); e != nil {
		return nil, e
	}

	return j, nil
}
