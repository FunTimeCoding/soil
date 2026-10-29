package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListPrefixes(x context.Context) ([]*client.Prefix, error) {
	r, e := c.client.ListPrefixesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list prefixes", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
