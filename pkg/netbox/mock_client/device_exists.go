package mock_client

func (c *Client) DeviceExists(_ string) (bool, error) {
	return false, c.Failure
}
