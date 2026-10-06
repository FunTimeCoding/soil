package loki

import "time"

func (c *Client) Labels(
	start time.Time,
	end time.Time,
) ([]string, error) {
	return c.basic.Labels(start, end)
}
