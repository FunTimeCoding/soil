package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteSnippet(
	x context.Context,
	name string,
	p *client.DeleteSnippetParams,
) (*client.StatusResult, error) {
	r, e := c.client.DeleteSnippetWithResponse(x, name, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("delete snippet", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
