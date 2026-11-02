package client

import "time"

func (c *Client) WithTimeout(d time.Duration) *Client {
	c.pool.Client().Timeout = d

	return c
}
