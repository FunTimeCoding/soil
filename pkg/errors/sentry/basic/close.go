package basic

func (c *Client) Close() {
	c.requester.Close()
}
