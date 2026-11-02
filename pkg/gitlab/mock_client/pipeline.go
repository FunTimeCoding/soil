package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline_detail"
)

func (c *Client) Pipeline(
	_ context.Context,
	_ int64,
	_ int64,
) (*pipeline_detail.Detail, error) {
	return nil, nil
}
