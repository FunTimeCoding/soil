package chip

func (c *Chip) Kind(kinds ...string) *Chip {
	c.kinds = append(c.kinds, kinds...)

	return c
}
