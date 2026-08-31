package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ToggleChecklistItem(
	key string,
	index int,
) *response.Response {
	result, e := c.client.ToggleChecklistItem(c.context, key, index)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
