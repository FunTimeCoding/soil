package mock_client

func (c *Client) Runbook(_ string) (string, error) {
	return "", c.failure
}
