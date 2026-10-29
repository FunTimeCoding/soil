package mock_client

func (c *Client) DeleteCertificateRequest(_ string, _ string) error {
	return c.Failure
}
