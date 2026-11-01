package mock_client

func (c *Client) PulseCruise(
	sessionIdentifier string,
	body string,
) (bool, error) {
	return c.recordPulse(sessionIdentifier, body)
}
