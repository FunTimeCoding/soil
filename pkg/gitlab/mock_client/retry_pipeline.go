package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"

func (c *Client) RetryPipeline(
	_ int64,
	_ int64,
) (*pipeline_detail.Detail, error) {
	return nil, nil
}
