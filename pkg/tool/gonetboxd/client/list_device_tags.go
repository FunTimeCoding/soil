package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListDeviceTags(
	x context.Context,
	device string,
) ([]string, error) {
	r, e := c.client.ListDeviceTagsWithResponse(x, device)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list device tags", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
