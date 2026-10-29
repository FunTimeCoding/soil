package openai

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/openai/refusal"
)

func (c *Client) Reachable(x context.Context) error {
	_, e := c.client.Models.List(x)

	return refusal.Classify(e)
}
