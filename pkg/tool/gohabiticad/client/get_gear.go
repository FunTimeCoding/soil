package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetGear(x context.Context) (*client.GearResult, error) {
	r, e := c.client.GetGearWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get gear", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
