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
	status string,
) string {
	body := client.CreateAddressRequest{
		Interface: interfaceName,
		Address:   address,
	}

	if status != "" {
		body.Status = &status
	}

	result, e := c.client.CreateAddress(c.context, device, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
