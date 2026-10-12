package mock_client

func (c *Client) Calls() []string {
	return c.calls
}
