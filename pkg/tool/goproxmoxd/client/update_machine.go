package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) UpdateMachine(
	identifier int64,
	node *string,
	body client.UpdateMachineJSONRequestBody,
) string {
	result, e := c.client.UpdateMachineWithResponse(
		c.context,
		identifier,
		&client.UpdateMachineParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
