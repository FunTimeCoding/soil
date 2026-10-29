package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTopAlerts(
	x context.Context,
	p *client.GetTopAlertsParams,
) ([]client.TopAlertsResponse, error) {
	r, e := c.client.GetTopAlertsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("top alerts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
