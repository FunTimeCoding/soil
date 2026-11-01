package card

func (c *Card) Link(path string) *Card {
	c.link = path

	return c
}
