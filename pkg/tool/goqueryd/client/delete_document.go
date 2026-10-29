package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteDocument(
	x context.Context,
	p *client.DeleteDocumentParams,
) (bool, error) {
	r, e := c.client.DeleteDocumentWithResponse(x, p)

	if e != nil {
		return false, e
	}

	if r.JSON200 == nil {
		return false, refusal.New("delete document", r.HTTPResponse, r.Body)
	}

	return r.JSON200.Deleted != nil && *r.JSON200.Deleted, nil
}
