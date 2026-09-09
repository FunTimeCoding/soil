package mock_client

func (c *Client) CompareFiles(
	_ int64,
	_ string,
	_ string,
) ([]string, error) {
	return nil, nil
}
