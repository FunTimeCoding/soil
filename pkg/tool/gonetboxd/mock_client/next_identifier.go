package mock_client

func (c *Client) nextIdentifier() int32 {
	c.identifier++

	return c.identifier
}
