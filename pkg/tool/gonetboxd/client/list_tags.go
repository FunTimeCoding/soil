package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListTags() string {
	result, e := c.client.ListTags(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
