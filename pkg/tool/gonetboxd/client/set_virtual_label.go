package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) SetVirtualLabel(
	x context.Context,
	machine string,
	key string,
	value string,
) (client.Label, error) {
	r, e := c.client.SetVirtualLabelWithResponse(
		x,
		machine,
		key,
		client.SetVirtualLabelJSONRequestBody{Value: value},
	)

	if e != nil {
		return client.Label{}, e
	}

	if r.JSON200 == nil {
		return client.Label{}, refusal.New(
			"set virtual label",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
