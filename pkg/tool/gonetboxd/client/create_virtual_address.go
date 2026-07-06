package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualAddress(
	vmName string,
	interfaceName string,
	address string,
) string {
	result, e := c.client.CreateVirtualAddress(
		c.context,
		vmName,
		client.CreateAddressRequest{
			Interface: interfaceName,
			Address:   address,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
