package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreateWirelessNetwork(_ context.Context, _ string) (client.WirelessNetwork, error) {
	return client.WirelessNetwork{}, c.Failure
}
