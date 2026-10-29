package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTag(
	x context.Context,
	p *client.GetTagParams,
) (*client.SourceTypeTag, error) {
	r, e := c.client.GetTagWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get tag", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
