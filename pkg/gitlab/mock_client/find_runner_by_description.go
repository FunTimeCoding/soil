package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) FindRunnerByDescription(
	_ context.Context,
	_ string,
) (*runner.Runner, bool, error) {
	return nil, false, c.failure
}
