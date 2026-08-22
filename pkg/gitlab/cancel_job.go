package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CancelJob(
	project int64,
	job int64,
) (*gitlab.Job, error) {
	result, _, e := c.client.Jobs.CancelJob(project, job)

	return result, wrapError(e)
}
