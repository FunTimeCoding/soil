package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline"

func (c *Client) Pipelines(
	_ int64,
	_ string,
	_ string,
	_ int64,
) ([]*pipeline.Pipeline, error) {
	return nil, nil
}
