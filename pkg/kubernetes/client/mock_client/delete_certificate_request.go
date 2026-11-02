package mock_client

func (c *Client) DeleteCertificateRequest(
	_ string,
	_ string,
) error {
	if c.Failure != nil {
		return c.Failure
	}

	return c.failures.Caller()
}
