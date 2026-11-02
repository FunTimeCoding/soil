package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostEvent(
	x context.Context,
	b client.PostEventJSONRequestBody,
) (*client.EventResponse, error) {
	r, e := c.client.PostEventWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post event", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
