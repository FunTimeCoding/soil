package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) MustProjectJobs(p *project.Project) []*job.Job {
	result, e := c.ProjectJobs(p)
	errors.PanicOnError(e)

	return result
}
