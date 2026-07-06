package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetTransitions(key string) string {
	result, e := c.client.GetTransitions(c.context, key)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
