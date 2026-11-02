package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateSnippet(
	x context.Context,
	p *client.CreateSnippetParams,
	b client.CreateSnippetJSONRequestBody,
) (*client.SnippetCreateResult, error) {
	r, e := c.client.CreateSnippetWithResponse(x, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create snippet", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
