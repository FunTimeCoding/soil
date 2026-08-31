package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ListVirtualAddresses(name string) *response.Response {
	result, e := c.client.ListVirtualAddresses(c.context, name)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
