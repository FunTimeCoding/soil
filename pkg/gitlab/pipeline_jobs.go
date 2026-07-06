package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) PipelineJobs(
	project int64,
	pipeline int64,
) ([]*job.Job, error) {
	result, _, e := c.client.Jobs.ListPipelineJobs(
		project,
		pipeline,
		&gitlab.ListJobsOptions{IncludeRetried: new(true)},
	)

	if e != nil {
		return nil, e
	}

	return c.enrichJobs(job.NewSlice(result)), nil
}
