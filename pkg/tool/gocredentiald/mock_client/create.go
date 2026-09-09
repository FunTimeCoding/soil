package mock_client

func (c *Client) Create(
	_ string,
	_ string,
	_ map[string]string,
) (string, error) {
	return "", nil
}
