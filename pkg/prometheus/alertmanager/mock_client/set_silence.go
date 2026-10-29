package mock_client

import "time"

func (c *Client) SetSilence(
	_ string,
	_ string,
	_ time.Duration,
) (string, error) {
	return "", c.failure
}
