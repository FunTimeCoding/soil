package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListNodes() string {
	result, e := c.client.ListNodesWithResponse(
		c.context,
		&client.ListNodesParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
