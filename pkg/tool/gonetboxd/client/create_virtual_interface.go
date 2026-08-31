package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualInterface(
	vmName string,
	name string,
) *response.Response {
	result, e := c.client.CreateVirtualInterface(
		c.context,
		vmName,
		client.CreateVirtualInterfaceRequest{Name: name},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
