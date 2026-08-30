package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteComment(
	key string,
	identifier string,
) (string, int) {
	result, e := c.client.DeleteComment(c.context, key, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
