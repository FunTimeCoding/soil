package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetEvents(
	x context.Context,
	p *client.GetEventsParams,
) ([]client.EventEntry, error) {
	r, e := c.client.GetEventsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get events", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
