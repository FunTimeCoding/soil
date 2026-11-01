package card

func (c *Card) Terms(text string) *Card {
	c.terms = text

	return c
}
