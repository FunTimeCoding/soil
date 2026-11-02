package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateTag(
	x context.Context,
	b client.CreateTagJSONRequestBody,
) (*client.Tag, error) {
	r, e := c.client.CreateTagWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("create tag", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
