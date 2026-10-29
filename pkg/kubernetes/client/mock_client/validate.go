package mock_client

func (c *Client) Validate() error {
	return c.Failure
}
