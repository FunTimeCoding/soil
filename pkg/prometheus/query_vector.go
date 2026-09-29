package prometheus

import "time"

func (c *Client) QueryVector(q string) (float64, error) {
	return c.QueryFloat(q, time.Now().Add(-time.Hour))
}
