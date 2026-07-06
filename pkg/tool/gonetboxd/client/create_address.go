package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateAddress(
	device string,
	interfaceName string,
	address string,
) string {
	result, e := c.client.CreateAddress(
		c.context,
		device,
		client.CreateAddressRequest{
			Interface: interfaceName,
			Address:   address,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
