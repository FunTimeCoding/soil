package mock_client

func (c *Client) AllMetrics() ([]string, error) {
	return nil, c.Failure
}
