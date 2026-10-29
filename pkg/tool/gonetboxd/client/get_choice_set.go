package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetChoiceSet(
	x context.Context,
	name string,
) (client.ChoiceSet, error) {
	r, e := c.client.GetChoiceSetWithResponse(x, name)

	if e != nil {
		return client.ChoiceSet{}, e
	}

	if r.JSON200 == nil {
		return client.ChoiceSet{}, refusal.New(
			"get choice set",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
