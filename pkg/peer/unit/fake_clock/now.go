package fake_clock

import "time"

func (c *Clock) Now() time.Time {
	return c.now
}
