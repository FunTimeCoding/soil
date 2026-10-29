package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateChoiceSet(
	x context.Context,
	name string,
	choices []string,
) (client.ChoiceSet, error) {
	r, e := c.client.UpdateChoiceSetWithResponse(
		x,
		name,
		client.UpdateChoiceSetRequest{Choices: choices},
	)

	if e != nil {
		return client.ChoiceSet{}, e
	}

	if r.JSON200 == nil {
		return client.ChoiceSet{}, refusal.New(
			"update choice set",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
