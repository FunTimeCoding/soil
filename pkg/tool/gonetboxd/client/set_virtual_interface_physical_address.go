package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SetVirtualInterfacePhysicalAddress(
	machine string,
	name string,
	address string,
) *response.Response {
	result, e := c.client.SetVirtualInterfacePhysicalAddress(
		c.context,
		machine,
		name,
		client.PhysicalAddressRequest{Address: address},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
