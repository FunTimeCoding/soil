package mock_client

func (c *Client) Query(_ string, _ any) error {
	return c.failure
}
