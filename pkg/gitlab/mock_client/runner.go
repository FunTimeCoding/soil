package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) Runner(
	_ context.Context,
	_ int64,
) (*runner.Runner, error) {
	return nil, c.failure
}
