package mock_client

func (c *Client) Keys(identifier string) []string {
	return c.keys[identifier]
}
