package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) PipelineJob(
	project int64,
	job int64,
) (*gitlab.Job, error) {
	result, _, e := c.client.Jobs.GetJob(project, job)

	return result, wrapError(e)
}
