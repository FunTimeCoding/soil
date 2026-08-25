package mock_client

func (c *Client) SetFailure(e error) {
	c.failure = e
}
