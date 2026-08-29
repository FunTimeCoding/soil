package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteHost(
	identifier string,
	apply *bool,
) string {
	result, e := c.client.DeleteHost(
		c.context,
		identifier,
		&client.DeleteHostParams{Apply: apply},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
