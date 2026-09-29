package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) FailedJobs() ([]*job.Job, error) {
	var result []*job.Job
	jobs, e := c.Jobs()

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
