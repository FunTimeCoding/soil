package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualDisk(
	machine string,
	name string,
	size int32,
) *response.Response {
	result, e := c.client.CreateVirtualDisk(
		c.context,
		machine,
		client.CreateVirtualDiskRequest{Name: name, Size: size},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
