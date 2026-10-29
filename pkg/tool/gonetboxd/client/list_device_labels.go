package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListDeviceLabels(
	x context.Context,
	device string,
) ([]*client.Label, error) {
	r, e := c.client.ListDeviceLabelsWithResponse(x, device)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list device labels", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
