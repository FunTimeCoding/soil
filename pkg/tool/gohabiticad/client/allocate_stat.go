package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AllocateStat(
	x context.Context,
	stat string,
) (*client.Stats, error) {
	r, e := c.client.AllocateStatWithResponse(
		x,
		client.AllocateStatParamsStat(stat),
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("allocate stat", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
