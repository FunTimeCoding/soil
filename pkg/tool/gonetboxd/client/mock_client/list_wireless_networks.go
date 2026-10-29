package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListWirelessNetworks(_ context.Context) ([]*client.WirelessNetwork, error) {
	return nil, c.Failure
}
