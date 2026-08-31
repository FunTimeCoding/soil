package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) RemoveVirtualLabel(
	machine string,
	key string,
) *response.Response {
	result, e := c.client.RemoveVirtualLabel(c.context, machine, key)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
