package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) enrichJobs(v []*job.Job) ([]*job.Job, error) {
	for _, j := range v {
		if _, e := c.enrichJob(j); e != nil {
			return nil, e
		}
	}

	return v, nil
}
