package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) enrichJobs(
	x context.Context,
	v []*job.Job,
) ([]*job.Job, error) {
	for _, j := range v {
		if _, e := c.enrichJob(x, j); e != nil {
			return nil, e
		}
	}

	return v, nil
}
