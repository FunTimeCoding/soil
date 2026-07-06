package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListNetworks(name string) string {
	result, e := c.client.ListNetworksWithResponse(
		c.context,
		name,
		&client.ListNetworksParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
