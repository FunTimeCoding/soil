package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) TagRunners(
	_ context.Context,
	_ string,
) ([]*runner.Runner, error) {
	return nil, c.failure
}
