package telemetry

import "time"

func (c *Client) WithTimeout(d time.Duration) *Client {
	c.client.WithTimeout(d)

	return c
}
