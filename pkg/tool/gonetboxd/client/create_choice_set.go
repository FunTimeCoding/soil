package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateChoiceSet(
	x context.Context,
	name string,
	choices []string,
) (client.ChoiceSet, error) {
	r, e := c.client.CreateChoiceSetWithResponse(
		x,
		client.CreateChoiceSetRequest{Name: name, Choices: choices},
	)

	if e != nil {
		return client.ChoiceSet{}, e
	}

	if r.JSON201 == nil {
		return client.ChoiceSet{}, refusal.New(
			"create choice set",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
