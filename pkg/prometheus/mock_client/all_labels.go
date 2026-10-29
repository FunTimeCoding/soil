package mock_client

func (c *Client) AllLabels() ([]string, error) {
	return nil, c.Failure
}
