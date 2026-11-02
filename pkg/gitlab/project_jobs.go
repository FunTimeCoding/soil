package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ProjectJobs(
	x context.Context,
	p *project.Project,
) ([]*job.Job, error) {
	result, _, e := c.client.Jobs.ListProjectJobs(
		p.Identifier,
		nil,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichProjectJobs(x, job.NewSlice(result), p)
}
