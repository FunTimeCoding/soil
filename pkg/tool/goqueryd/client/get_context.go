package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetContext(x context.Context) ([]*client.ContextEntry, error) {
	r, e := c.client.GetContextWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list contexts", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
