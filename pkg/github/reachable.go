package github

import (
	"context"
	"github.com/funtimecoding/soil/pkg/github/refusal"
)

func (c *Client) Reachable(x context.Context) error {
	_, _, e := c.client.RateLimit.Get(x)

	return refusal.Classify(e)
}
