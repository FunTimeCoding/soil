package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTransitions(
	x context.Context,
	key string,
) ([]*client.JiraTransition, error) {
	r, e := c.client.GetTransitionsWithResponse(x, key)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("transitions", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
