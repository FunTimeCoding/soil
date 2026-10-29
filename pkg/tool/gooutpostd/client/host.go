package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Host(x context.Context) (*client.Host, error) {
	r, e := c.client.GetHostWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("host", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
