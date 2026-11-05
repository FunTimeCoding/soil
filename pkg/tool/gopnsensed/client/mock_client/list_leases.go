package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) ListLeases(
	_ context.Context,
	_ string,
) ([]client.Lease, error) {
	return c.LeaseList, c.Failure
}
