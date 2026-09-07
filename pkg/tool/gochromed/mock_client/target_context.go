package mock_client

import "context"

func (c *Client) TargetContext(identifier string) context.Context {
	return c.AcquireTarget(identifier)
}
