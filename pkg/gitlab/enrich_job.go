package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) enrichJob(j *job.Job) (*job.Job, error) {
	identifier := j.ProjectIdentifier()

	if identifier == 0 {
		return nil, validation.New("job %d has no project", j.Identifier)
	}

	p, e := c.Project(identifier)

	if e != nil {
		return nil, e
	}

	j.Project = p

	if f := c.enrichJobCommon(j); f != nil {
		return nil, f
	}

	return j, nil
}
