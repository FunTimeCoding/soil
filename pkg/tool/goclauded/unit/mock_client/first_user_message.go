package mock_client

func (c *Client) FirstUserMessage(sessionIdentifier string) string {
	return c.FirstMessages[sessionIdentifier]
}
