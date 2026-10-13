package client_tester

func (c *Client) Close() {
	c.Closed++
}
