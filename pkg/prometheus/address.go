package prometheus

func (c *Client) Address() string {
	return c.link.Address()
}
