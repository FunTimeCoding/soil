package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListStorageContent(
	name string,
	storage string,
) string {
	result, e := c.client.ListStorageContentWithResponse(
		c.context,
		name,
		storage,
		&client.ListStorageContentParams{
			Instance: &c.instance,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
