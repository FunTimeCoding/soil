package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeriveHardwareAddress(identifier int) string {
	result, e := c.client.DeriveHardwareAddressWithResponse(
		c.context,
		&client.DeriveHardwareAddressParams{
			Instance:   &c.instance,
			Identifier: identifier,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
