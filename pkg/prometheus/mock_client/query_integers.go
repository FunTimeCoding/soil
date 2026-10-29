package mock_client

import "time"

func (c *Client) QueryIntegers(
	_ string,
	_ time.Time,
) (map[string]int, error) {
	return nil, c.Failure
}
