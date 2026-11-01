package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/reachable_probe"
)

func (c *Client) Reachable(x context.Context) error {
	return reachable_probe.Reach(x, c.pool.Client(), c.base, c.token)
}
