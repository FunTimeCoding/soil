package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/response"
)

func (c *Client) GraphRunner(
	_ context.Context,
	_ int64,
) (*response.Runner, error) {
	return nil, c.failure
}
