package chip

func (c *Chip) Current() *Chip {
	c.current = true
	c.active = true

	return c
}
