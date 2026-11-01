package card

func (c *Card) Meta(lines ...string) *Card {
	c.meta = append(c.meta, lines...)

	return c
}
