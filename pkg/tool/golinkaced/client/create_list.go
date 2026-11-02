package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateList(
	x context.Context,
	b client.CreateListJSONRequestBody,
) (*client.List, error) {
	r, e := c.client.CreateListWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create list", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
