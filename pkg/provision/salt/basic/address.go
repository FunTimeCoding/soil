package basic

func (c *Client) Address() string {
	return c.base.String()
}
