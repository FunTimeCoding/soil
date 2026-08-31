package chart

import "time"

func (c *Chart) WithNow(now time.Time) *Chart {
	c.now = now

	return c
}
