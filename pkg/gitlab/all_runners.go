package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) allRunners(
	x context.Context,
	o *gitlab.ListRunnersOptions,
) ([]*gitlab.Runner, *gitlab.Response, error) {
	return c.client.Runners.ListAllRunners(o, gitlab.WithContext(x))
}
