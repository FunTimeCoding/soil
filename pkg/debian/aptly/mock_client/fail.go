package mock_client

func (c *Client) Fail(e error) {
	c.fail = e
}
