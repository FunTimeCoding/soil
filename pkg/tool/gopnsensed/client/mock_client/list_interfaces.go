package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) ListInterfaces(
	_ context.Context,
) ([]client.NetworkInterface, error) {
	return c.InterfaceList, c.Failure
}
