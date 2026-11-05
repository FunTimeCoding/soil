package salt

import "context"

func (c *Client) Reachable(x context.Context) error {
	return c.Login(x)
}
