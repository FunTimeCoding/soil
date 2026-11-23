package graph

func (c *Client) Close() {
	c.requester.Close()
	c.pool.Close()
}
