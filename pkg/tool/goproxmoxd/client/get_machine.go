package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetMachine(
	identifier int64,
	node *string,
) string {
	result, e := c.client.GetMachineWithResponse(
		c.context,
		identifier,
		&client.GetMachineParams{Instance: &c.instance, Node: node},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
