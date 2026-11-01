package chip

func (c *Chip) Color(color string) *Chip {
	c.color = color

	return c
}
