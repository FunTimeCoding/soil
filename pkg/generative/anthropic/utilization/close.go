package utilization

func (c *Client) Close() {
	c.requester.Close()
}
