package mock_client

func (c *Client) PulseImmediate(
	sessionIdentifier string,
	body string,
) (bool, error) {
	return c.recordPulse(sessionIdentifier, body)
}
