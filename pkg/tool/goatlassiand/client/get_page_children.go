package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetPageChildren(identifier string) (string, int) {
	result, e := c.client.GetPageChildren(c.context, identifier)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
