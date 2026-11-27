package chromium

import "context"

func (c *Client) Context() context.Context {
	return c.context
}
