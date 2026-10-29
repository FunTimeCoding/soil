package mock_client

func (c *Client) Forget() {
	c.Forgotten++
}
