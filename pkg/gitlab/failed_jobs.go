package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) FailedJobs() []*job.Job {
	var result []*job.Job

	for _, j := range c.Jobs() {
		if j.Fail() {
			result = append(result, j)
		}
	}

	return result
}
