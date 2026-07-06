package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListProjects() string {
	result, e := c.client.ListProjects(c.context)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
