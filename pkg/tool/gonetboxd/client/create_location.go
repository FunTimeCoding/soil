package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateLocation(
	x context.Context,
	name string,
	site string,
) (client.Location, error) {
	r, e := c.client.CreateLocationWithResponse(
		x,
		client.CreateLocationRequest{Name: name, Site: site},
	)

	if e != nil {
		return client.Location{}, e
	}

	if r.JSON201 == nil {
		return client.Location{}, refusal.New(
			"create location",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
