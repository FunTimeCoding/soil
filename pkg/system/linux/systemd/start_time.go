package systemd

import "time"

func (c *Client) StartTime(name string) time.Time {
	return c.Show(name).ActiveEnter
}
