package mock_client

func (c *Client) DeleteNote(_ int) error {
	return c.Failure
}
