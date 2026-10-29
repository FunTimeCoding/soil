package mock_client

import "time"

func (c *Client) QueryFloat(
	q string,
	_ time.Time,
) (float64, error) {
	if c.Failure != nil {
		return 0, c.Failure
	}

	return c.floats[q], nil
}
