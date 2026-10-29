package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetAlertsByIdentifier(
	x context.Context,
	p *client.GetAlertsByIdentifierParams,
) ([]client.AlertsResponse, error) {
	r, e := c.client.GetAlertsByIdentifierWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("alerts by identifier", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
