package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteChecklistItem(
	key string,
	index int,
) (string, int) {
	result, e := c.client.DeleteChecklistItem(c.context, key, index)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
