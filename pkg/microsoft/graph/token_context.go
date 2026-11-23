package graph

import (
	"context"
	"golang.org/x/oauth2"
)

func (c *Client) tokenContext() context.Context {
	return context.WithValue(
		context.Background(),
		oauth2.HTTPClient,
		c.pool.Client(),
	)
}
