package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) DeleteComment(
	key string,
	identifier string,
) *response.Response {
	result, e := c.client.DeleteComment(c.context, key, identifier)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
