package mock_client

func (c *Client) Close() {
	c.Closed++
}
