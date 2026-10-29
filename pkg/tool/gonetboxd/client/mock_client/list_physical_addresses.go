package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListPhysicalAddresses(
	_ context.Context,
) ([]client.PhysicalAddressOwner, error) {
	return c.PhysicalAddressList, c.Failure
}
