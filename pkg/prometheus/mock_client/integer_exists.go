package mock_client

import "time"

func (c *Client) IntegerExists(
	_ string,
	_ time.Time,
) (bool, error) {
	return false, c.Failure
}
