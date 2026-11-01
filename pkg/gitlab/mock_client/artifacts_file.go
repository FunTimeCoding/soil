package mock_client

func (c *Client) ArtifactsFile(_ int64, _ string, _ string) (string, error) {
	return "", c.failure
}
