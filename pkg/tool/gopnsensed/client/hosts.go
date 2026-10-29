package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Hosts(
	x context.Context,
	phrase string,
) ([]client.Host, error) {
	r, e := c.client.ListHostsWithResponse(
		x,
		&client.ListHostsParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("hosts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
