package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) RunnersByTag(
	_ context.Context,
	_ string,
) ([]*runner.Runner, error) {
	return nil, c.failure
}
