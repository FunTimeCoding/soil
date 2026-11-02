package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RemoveList(
	x context.Context,
	identifier int32,
	b client.RemoveListJSONRequestBody,
) (*client.Link, error) {
	r, e := c.client.RemoveListWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("remove list", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
