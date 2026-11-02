package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetLinks(
	x context.Context,
	p *client.GetLinksParams,
) (*client.LinkPage, error) {
	r, e := c.client.GetLinksWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get links", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
