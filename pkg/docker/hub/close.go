package hub

func (c *Client) Close() {
	c.requester.Close()
}
