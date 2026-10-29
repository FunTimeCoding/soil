package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetDocument(
	x context.Context,
	p *client.GetDocumentParams,
) (*client.SearchResult, error) {
	r, e := c.client.GetDocumentWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get document", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
