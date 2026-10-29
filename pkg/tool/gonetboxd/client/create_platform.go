package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePlatform(
	x context.Context,
	name string,
) (client.Platform, error) {
	r, e := c.client.CreatePlatformWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.Platform{}, e
	}

	if r.JSON201 == nil {
		return client.Platform{}, refusal.New(
			"create platform",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
