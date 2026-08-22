package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Pipeline(
	project int64,
	pipeline int64,
) (*gitlab.Pipeline, error) {
	result, _, e := c.client.Pipelines.GetPipeline(project, pipeline)

	return result, wrapError(e)
}
