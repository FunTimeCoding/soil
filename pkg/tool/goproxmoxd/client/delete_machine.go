package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteMachine(
	identifier int64,
	node *string,
	purge *bool,
) string {
	result, e := c.client.DeleteMachineWithResponse(
		c.context,
		identifier,
		&client.DeleteMachineParams{
			Instance: &c.instance,
			Node:     node,
			Purge:    purge,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
