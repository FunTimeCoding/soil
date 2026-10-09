package unit

import "time"

func (c *clock) Now() time.Time {
	return c.now
}
