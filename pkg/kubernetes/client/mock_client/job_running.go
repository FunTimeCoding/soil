package mock_client

func (c *Client) JobRunning(
	_ string,
	_ string,
) (bool, error) {
	return false, c.Failure
}
