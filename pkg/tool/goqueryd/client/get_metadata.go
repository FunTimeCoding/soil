package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetMetadata(
	x context.Context,
	p *client.GetMetadataParams,
) ([]*client.Facet, error) {
	r, e := c.client.GetMetadataWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list metadata", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
