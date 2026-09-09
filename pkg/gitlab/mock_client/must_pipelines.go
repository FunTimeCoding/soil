package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/pipeline"

func (c *Client) MustPipelines(_ int64) []*pipeline.Pipeline {
	return nil
}
