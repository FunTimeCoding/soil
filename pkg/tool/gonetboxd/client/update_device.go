package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateDevice(
	name string,
	newName string,
	primaryAddress string,
) string {
	body := client.UpdateDeviceRequest{}

	if newName != "" {
		body.Name = &newName
	}

	if primaryAddress != "" {
		body.PrimaryAddress = &primaryAddress
	}

	result, e := c.client.UpdateDevice(c.context, name, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
