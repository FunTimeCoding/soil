package mock_client

func (c *Client) Packages(string) ([]string, error) {
	return nil, c.fail
}
