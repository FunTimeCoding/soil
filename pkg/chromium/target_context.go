package chromium

import "context"

func (c *Client) TargetContext(identifier string) (context.Context, error) {
	return c.AcquireTarget(identifier)
}
