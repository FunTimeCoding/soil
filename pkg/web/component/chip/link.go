package chip

func (c *Chip) Link(path string) *Chip {
	c.link = path

	return c
}
