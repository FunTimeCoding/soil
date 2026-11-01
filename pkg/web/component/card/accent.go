package card

func (c *Card) Accent(color string) *Card {
	c.accent = color

	return c
}
