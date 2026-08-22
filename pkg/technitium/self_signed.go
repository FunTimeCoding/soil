package technitium

func (c *Client) SelfSigned() *Client {
	c.basic.SelfSigned()

	return c
}
