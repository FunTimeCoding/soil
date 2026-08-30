package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateWirelessNetwork(ssid string) (string, int) {
	result, e := c.client.CreateWirelessNetwork(
		c.context,
		client.CreateNameRequest{Name: ssid},
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
