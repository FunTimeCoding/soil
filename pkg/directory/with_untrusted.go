package directory

func (c *Client) WithUntrusted() *Client {
	c.untrusted = true

	return c
}
