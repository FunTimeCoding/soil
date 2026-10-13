package aptly

func (c *Client) Close() {
	c.requester.Close()
}
