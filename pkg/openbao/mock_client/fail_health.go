package mock_client

func (c *Client) FailHealth(e error) {
	c.health = e
}
