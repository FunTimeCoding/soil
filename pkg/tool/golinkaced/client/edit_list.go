package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditList(
	x context.Context,
	identifier int32,
	b client.EditListJSONRequestBody,
) (*client.List, error) {
	r, e := c.client.EditListWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit list", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
