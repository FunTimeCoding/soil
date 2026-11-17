package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) SetInterfaceEthernetPower(
	_ context.Context,
	_ string,
	_ string,
	mode string,
	kind string,
) (client.Interface, error) {
	c.EthernetPower = &client.EthernetPowerRequest{Mode: mode}

	if kind != "" {
		c.EthernetPower.Type = &kind
	}

	return client.Interface{}, c.Failure
}
