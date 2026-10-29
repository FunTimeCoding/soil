package mock_client

func (c *Client) ExecuteSimple(_ string, _ string, _ ...string) (string, string, error) {
	return "", "", c.Failure
}
