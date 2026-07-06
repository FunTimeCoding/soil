package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DownloadLocator(
	node string,
	storage string,
	body client.DownloadLocatorJSONRequestBody,
) string {
	result, e := c.client.DownloadLocatorWithResponse(
		c.context,
		node,
		storage,
		&client.DownloadLocatorParams{
			Instance: &c.instance,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
