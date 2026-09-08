package technitium

func (c *Client) Untrusted() *Client {
	c.basic.Untrusted()

	return c
}
