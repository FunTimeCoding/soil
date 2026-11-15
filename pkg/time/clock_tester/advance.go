package clock_tester

import "time"

func (c *Clock) Advance(d time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.now = c.now.Add(d)
}
