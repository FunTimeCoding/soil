package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"

func (c *Client) RetryPipeline(
	project int64,
	identifier int64,
) (*pipeline_detail.Detail, error) {
	result, _, e := c.client.Pipelines.RetryPipelineBuild(project, identifier)

	if e != nil {
		return nil, wrapError(e)
	}

	return pipeline_detail.New(result), nil
}
