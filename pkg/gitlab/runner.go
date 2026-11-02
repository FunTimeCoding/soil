package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Runner(
	x context.Context,
	identifier int64,
) (*runner.Runner, error) {
	result, _, e := c.client.Runners.GetRunnerDetails(
		identifier,
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichRunner(x, runner.NewDetail(result))
}
