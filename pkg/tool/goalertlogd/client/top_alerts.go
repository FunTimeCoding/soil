package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"time"
)

func (c *Client) TopAlerts(
	n int,
	start time.Time,
	end time.Time,
) *response.Response {
	result, e := c.client.GetTopAlerts(
		c.context,
		&client.GetTopAlertsParams{N: &n, Start: &start, End: &end},
	)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
