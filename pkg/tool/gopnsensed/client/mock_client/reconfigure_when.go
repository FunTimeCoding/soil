package mock_client

func (c *Client) reconfigureWhen(apply bool) {
	if apply {
		c.Reconfigured++
	}
}
