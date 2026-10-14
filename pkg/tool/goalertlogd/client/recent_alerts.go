package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"time"
)

func (c *Client) RecentAlerts(
	start *time.Time,
	end *time.Time,
) *response.Response {
	result, e := c.client.GetRecentAlerts(
		c.context,
		&client.GetRecentAlertsParams{Start: start, End: end},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
