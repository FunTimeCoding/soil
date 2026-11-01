package mock_client

func (c *Client) Pulse(
	sessionIdentifier string,
	body string,
) (bool, error) {
	return c.recordPulse(sessionIdentifier, body)
}
