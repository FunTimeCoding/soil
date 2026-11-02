package hub

import "time"

func (c *Client) WithBackoff(d time.Duration) *Client {
	c.requester.WithBackoff(d)

	return c
}
