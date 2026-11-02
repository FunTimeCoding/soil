package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetSnippet(
	x context.Context,
	name string,
	p *client.GetSnippetParams,
) (*client.SnippetDetail, error) {
	r, e := c.client.GetSnippetWithResponse(x, name, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get snippet", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
