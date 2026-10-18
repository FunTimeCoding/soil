package fake_clock

import "time"

func (c *Clock) Advance(d time.Duration) {
	c.now = c.now.Add(d)
}
