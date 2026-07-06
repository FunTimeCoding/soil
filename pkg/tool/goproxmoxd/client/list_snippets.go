package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListSnippets() string {
	result, e := c.client.ListSnippetsWithResponse(
		c.context,
		&client.ListSnippetsParams{
			Instance: &c.instance,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
