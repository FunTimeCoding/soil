package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Leases(
	x context.Context,
	phrase string,
) ([]client.Lease, error) {
	r, e := c.client.ListLeasesWithResponse(
		x,
		&client.ListLeasesParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("leases", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
