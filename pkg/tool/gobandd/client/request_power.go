package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RequestPower(
	x context.Context,
	p client.PowerRequest,
) (*client.PowerResult, error) {
	r, e := c.client.RequestPowerWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("power", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
