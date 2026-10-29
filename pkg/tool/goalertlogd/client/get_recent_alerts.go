package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetRecentAlerts(
	x context.Context,
	p *client.GetRecentAlertsParams,
) ([]client.AlertsResponse, error) {
	r, e := c.client.GetRecentAlertsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("recent alerts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
