package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Pools(
	_ context.Context,
	_ string,
) ([]client.Pool, error) {
	return c.PoolList, c.Failure
}
