package mock_client

func (c *Client) Resolved(thread string) bool {
	return c.resolved[thread]
}
