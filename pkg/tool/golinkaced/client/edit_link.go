package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditLink(
	x context.Context,
	identifier int32,
	b client.EditLinkJSONRequestBody,
) (*client.Link, error) {
	r, e := c.client.EditLinkWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit link", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
