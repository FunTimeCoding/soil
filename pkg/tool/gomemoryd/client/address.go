package client

func (c *Client) Address() string {
	return c.base.String()
}
