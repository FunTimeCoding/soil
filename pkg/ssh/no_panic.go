package ssh

func (c *Client) NoPanic() *Client {
	c.Panic = false

	return c
}
