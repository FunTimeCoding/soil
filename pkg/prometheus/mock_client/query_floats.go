package mock_client

import "time"

func (c *Client) QueryFloats(
	_ string,
	_ time.Time,
) (map[string]float64, error) {
	return nil, c.Failure
}
