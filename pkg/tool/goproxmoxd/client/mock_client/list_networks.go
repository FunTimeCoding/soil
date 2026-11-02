package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListNetworks(
	_ context.Context,
	node string,
	_ *client.ListNetworksParams,
) ([]client.Network, error) {
	if e := c.node(node); e != nil {
		return nil, e
	}

	return c.NetworkList, nil
}
