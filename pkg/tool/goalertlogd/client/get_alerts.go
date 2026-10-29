package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetAlerts(
	x context.Context,
	p *client.GetAlertsParams,
) ([]client.AlertsResponse, error) {
	r, e := c.client.GetAlertsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("alerts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
