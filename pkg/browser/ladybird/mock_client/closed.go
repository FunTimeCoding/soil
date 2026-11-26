package mock_client

func (c *Client) Closed() bool {
	return c.closed.Load()
}
