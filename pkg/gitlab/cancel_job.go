package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) CancelJob(
	project int64,
	identifier int64,
) (*job.Job, error) {
	result, _, e := c.client.Jobs.CancelJob(project, identifier)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichJob(job.New(result)), nil
}
