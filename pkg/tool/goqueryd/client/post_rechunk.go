package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostRechunk(x context.Context) (*client.RechunkResult, error) {
	r, e := c.client.PostRechunkWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("rechunk", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
