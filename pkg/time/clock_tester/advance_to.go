package clock_tester

import "time"

func (c *Clock) AdvanceTo(instant time.Time) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.now = instant
}
