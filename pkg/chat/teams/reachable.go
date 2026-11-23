package teams

import "context"

func (c *Client) Reachable(x context.Context) error {
	return c.graph.Reachable(x)
}
