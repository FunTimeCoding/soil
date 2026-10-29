package mock_client

import "time"

func (c *Client) PostSilence(
	_ string,
	_ string,
	_ string,
	_ time.Time,
	_ time.Time,
) (string, error) {
	return "", c.failure
}
