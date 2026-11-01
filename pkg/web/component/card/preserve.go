package card

func (c *Card) Preserve() *Card {
	c.preserve = true

	return c
}
