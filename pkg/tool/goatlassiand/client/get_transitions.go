package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetTransitions(key string) *response.Response {
	result, e := c.client.GetTransitions(c.context, key)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
