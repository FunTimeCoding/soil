package mock_client

func (c *Client) Closed() int {
	return c.closed
}
