package mock_client

func (c *Client) TriggerCertificateIssuance(_ string, _ string) error {
	return c.Failure
}
