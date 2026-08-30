package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DownloadLocator(
	node string,
	storage string,
	body client.DownloadLocatorJSONRequestBody,
) (string, int) {
	result, e := c.client.DownloadLocatorWithResponse(
		c.context,
		node,
		storage,
		&client.DownloadLocatorParams{Instance: &c.instance},
		body,
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
