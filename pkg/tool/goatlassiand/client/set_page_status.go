package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) SetPageStatus(
	identifier string,
	status string,
) (string, int) {
	result, e := c.client.SetPageStatus(
		c.context,
		identifier,
		client.PageStatusRequest{Status: status},
	)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
