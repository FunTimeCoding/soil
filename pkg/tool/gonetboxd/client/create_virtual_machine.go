package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateVirtualMachine(
	name string,
	cluster string,
) string {
	result, e := c.client.CreateVirtualMachine(
		c.context,
		client.CreateVirtualMachineRequest{
			Name:    name,
			Cluster: cluster,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
