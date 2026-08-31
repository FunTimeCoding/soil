package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) GetChecklist(key string) *response.Response {
	result, e := c.client.GetChecklist(c.context, key)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
