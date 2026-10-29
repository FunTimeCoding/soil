package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Aliases(
	x context.Context,
	phrase string,
) ([]client.Alias, error) {
	r, e := c.client.ListAliasesWithResponse(
		x,
		&client.ListAliasesParams{Query: optional(phrase)},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("aliases", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
