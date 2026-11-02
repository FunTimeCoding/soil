package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AppendList(
	x context.Context,
	identifier int32,
	b client.AppendListJSONRequestBody,
) (*client.Link, error) {
	r, e := c.client.AppendListWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("append list", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
