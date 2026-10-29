package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Blocklists(
	x context.Context,
	phrase string,
) ([]client.Blocklist, error) {
	r, e := c.client.ListBlocklistsWithResponse(
		x,
		&client.ListBlocklistsParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("blocklists", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
