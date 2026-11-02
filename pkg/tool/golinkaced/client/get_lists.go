package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetLists(
	x context.Context,
	p *client.GetListsParams,
) (*client.ListPage, error) {
	r, e := c.client.GetListsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get lists", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
