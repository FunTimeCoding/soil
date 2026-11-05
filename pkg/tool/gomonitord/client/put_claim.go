package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PutClaim(
	x context.Context,
	item string,
	owner string,
) (*client.Claim, error) {
	r, e := c.client.PutClaimWithResponse(
		x,
		client.PutClaimJSONRequestBody{Item: item, Owner: owner},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("put claim", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
