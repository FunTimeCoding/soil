package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) StopMachine(
	identifier int64,
	node *string,
) string {
	result, e := c.client.StopMachineWithResponse(
		c.context,
		identifier,
		&client.StopMachineParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
