package card

func (c *Card) Subtitle(text string) *Card {
	c.subtitle = text

	return c
}
