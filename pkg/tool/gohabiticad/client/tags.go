package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Tags() string {
	result, e := c.client.GetTags(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
