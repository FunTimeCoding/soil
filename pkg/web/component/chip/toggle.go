package chip

func (c *Chip) Toggle(active bool) *Chip {
	c.toggle = true
	c.active = active

	return c
}
