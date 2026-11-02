package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) PipelineJobs(
	x context.Context,
	project int64,
	pipeline int64,
) ([]*job.Job, error) {
	result, _, e := c.client.Jobs.ListPipelineJobs(
		project,
		pipeline,
		&gitlab.ListJobsOptions{IncludeRetried: new(true)},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	p, f := c.Project(x, project)

	if f != nil {
		return nil, f
	}

	return c.enrichProjectJobs(x, job.NewSlice(result), p)
}
