package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetLinkTypes(x context.Context) ([]client.LinkType, error) {
	r, e := c.client.GetLinkTypesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("link types", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
