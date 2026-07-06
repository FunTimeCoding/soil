package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateSnippet(
	body client.CreateSnippetJSONRequestBody,
) string {
	result, e := c.client.CreateSnippetWithResponse(
		c.context,
		&client.CreateSnippetParams{
			Instance: &c.instance,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
