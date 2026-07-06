package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetSnippet(name string) string {
	result, e := c.client.GetSnippetWithResponse(
		c.context,
		name,
		&client.GetSnippetParams{
			Instance: &c.instance,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
