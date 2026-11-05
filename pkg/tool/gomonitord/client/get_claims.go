package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetClaims(x context.Context) ([]client.Claim, error) {
	r, e := c.client.GetClaimsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get claims", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
