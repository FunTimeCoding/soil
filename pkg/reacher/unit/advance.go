package unit

import "time"

func (c *clock) Advance(d time.Duration) {
	c.now = c.now.Add(d)
}
