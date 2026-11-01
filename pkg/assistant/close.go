package assistant

func (c *Client) Close() {
	c.hold.Close()
}
