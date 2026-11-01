package assistant

import "time"

func (c *Client) WithBackoff(
	minimum time.Duration,
	maximum time.Duration,
) *Client {
	c.hold.WithBackoff(minimum, maximum)

	return c
}
