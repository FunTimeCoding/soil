package mock_client

func (c *Client) RedactedMemories() map[int64]bool {
	return c.Redacted
}
