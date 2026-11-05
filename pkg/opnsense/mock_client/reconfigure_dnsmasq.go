package mock_client

func (c *Client) ReconfigureDnsmasq() error {
	return c.Failure
}
