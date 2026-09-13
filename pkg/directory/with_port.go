package directory

func (c *Client) WithPort(port int) *Client {
	c.port = port

	return c
}
