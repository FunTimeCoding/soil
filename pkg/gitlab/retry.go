package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) Retry(
	project int64,
	jobIdentifier int64,
) (*job.Job, error) {
	result, _, e := c.client.Jobs.RetryJob(project, jobIdentifier)

	if e != nil {
		return nil, e
	}

	return c.enrichJob(job.New(result)), nil
}
