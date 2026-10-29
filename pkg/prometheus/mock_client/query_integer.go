package mock_client

import "time"

func (c *Client) QueryInteger(
	_ string,
	_ time.Time,
) (int, error) {
	return 0, c.Failure
}
