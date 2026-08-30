package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteLink(identifier string) (string, int) {
	result, e := c.client.DeleteLink(c.context, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
