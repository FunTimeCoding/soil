package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddLink(
	x context.Context,
	b client.AddLinkJSONRequestBody,
) (*client.Link, error) {
	r, e := c.client.AddLinkWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("add link", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
