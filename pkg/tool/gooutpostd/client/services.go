package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) Services(
	x context.Context,
	origin string,
) ([]client.Service, error) {
	params := &client.ListServicesParams{}

	if origin != "" {
		params.Origin = &origin
	}

	r, e := c.client.ListServicesWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("services", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
