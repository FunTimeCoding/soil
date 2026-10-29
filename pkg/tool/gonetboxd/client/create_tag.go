package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTag(
	x context.Context,
	name string,
) (client.Tag, error) {
	r, e := c.client.CreateTagWithResponse(
		x,
		client.CreateNameRequest{Name: name},
	)

	if e != nil {
		return client.Tag{}, e
	}

	if r.JSON201 == nil {
		return client.Tag{}, refusal.New("create tag", r.HTTPResponse, r.Body)
	}

	return *r.JSON201, nil
}
