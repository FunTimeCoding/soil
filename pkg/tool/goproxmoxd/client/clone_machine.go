package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CloneMachine(
	identifier int64,
	node *string,
	body client.CloneMachineJSONRequestBody,
) string {
	result, e := c.client.CloneMachineWithResponse(
		c.context,
		identifier,
		&client.CloneMachineParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
