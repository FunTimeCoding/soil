package mock_client

func (c *Client) Trace(
	_ int64,
	_ int64,
) (string, error) {
	return "", nil
}
