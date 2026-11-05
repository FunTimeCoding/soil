package github

import (
	"github.com/funtimecoding/soil/pkg/github/job"
	"github.com/funtimecoding/soil/pkg/github/refusal"
)

func (c *Client) Jobs(
	owner string,
	name string,
	run int64,
) ([]*job.Job, error) {
	result, _, e := c.client.Actions.ListWorkflowJobs(
		c.context,
		owner,
		name,
		run,
		nil,
	)

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return job.NewSlice(result.Jobs), nil
}
