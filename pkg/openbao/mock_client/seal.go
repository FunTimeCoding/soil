package mock_client

func (c *Client) Seal() {
	c.status.Sealed = true
}
