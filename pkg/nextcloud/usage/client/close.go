package client

func (c *Client) Close() {
	c.requester.Close()
}
