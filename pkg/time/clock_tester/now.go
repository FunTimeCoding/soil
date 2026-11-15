package clock_tester

import "time"

func (c *Clock) Now() time.Time {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.now
}
