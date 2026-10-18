package openai

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, e := c.client.Models.List(x)

	return e
}
