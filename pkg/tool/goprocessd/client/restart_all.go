package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RestartAll(x context.Context) (int, error) {
	r, e := c.client.RestartAllWithResponse(x)

	if e != nil {
		return 0, e
	}

	if r.JSON200 == nil {
		return 0, refusal.New("restart all", r.HTTPResponse, r.Body)
	}

	return r.JSON200.Count, nil
}
