package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetPage(
	x context.Context,
	identifier string,
) (*client.ConfluencePageDetail, error) {
	r, e := c.client.GetPageWithResponse(x, identifier)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("page", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
