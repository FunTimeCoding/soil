package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Log(limit *int) *response.Response {
	result, e := c.client.FirewallLog(
		c.context,
		&client.FirewallLogParams{Limit: limit},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
