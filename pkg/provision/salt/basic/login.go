package basic

import "context"

func (c *Client) Login(x context.Context) error {
	return c.session.Login(x)
}
