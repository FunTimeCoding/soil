package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) AddVirtualTag(
	name string,
	tag string,
) *response.Response {
	result, e := c.client.AddVirtualTag(c.context, name, tag)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
