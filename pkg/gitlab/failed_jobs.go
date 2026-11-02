package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) FailedJobs(x context.Context) ([]*job.Job, error) {
	var result []*job.Job
	jobs, e := c.Jobs(x)

	if e != nil {
		return nil, e
	}

	for _, j := range jobs {
		if j.Fail() {
			result = append(result, j)
		}
	}

	return result, nil
}
