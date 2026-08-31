package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Alerts() *response.Response {
	result, e := c.client.GetAlerts(c.context, &client.GetAlertsParams{})
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
