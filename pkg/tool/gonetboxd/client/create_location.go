package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateLocation(
	name string,
	site string,
) string {
	result, e := c.client.CreateLocation(
		c.context,
		client.CreateLocationRequest{Name: name, Site: site},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
