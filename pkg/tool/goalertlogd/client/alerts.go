package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Alerts() string {
	result, e := c.client.GetAlerts(c.context, &client.GetAlertsParams{})
	errors.PanicOnError(e)

	return web.ReadString(result)
}
