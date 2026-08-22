package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ToggleChecklistItem(
	key string,
	index int,
) string {
	result, e := c.client.ToggleChecklistItem(c.context, key, index)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
