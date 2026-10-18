package salt

import "context"

func (c *Client) Login(x context.Context) error {
	return c.basic.Login(x)
}
