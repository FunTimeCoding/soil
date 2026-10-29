package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SetPageStatus(
	x context.Context,
	identifier string,
	status string,
) (*client.ConfluencePageDetail, error) {
	r, e := c.client.SetPageStatusWithResponse(
		x,
		identifier,
		client.PageStatusRequest{Status: status},
	)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("set page status", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
