package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Pools(
	x context.Context,
	phrase string,
) ([]client.Pool, error) {
	r, e := c.client.ListPoolsWithResponse(
		x,
		&client.ListPoolsParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("pools", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
