package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSnippets(
	x context.Context,
	p *client.ListSnippetsParams,
) ([]client.SnippetItem, error) {
	r, e := c.client.ListSnippetsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list snippets", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
