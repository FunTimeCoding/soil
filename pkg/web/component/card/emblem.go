package card

func (c *Card) Emblem(text string) *Card {
	c.emblem = text

	return c
}
