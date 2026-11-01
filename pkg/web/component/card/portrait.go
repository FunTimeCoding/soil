package card

func (c *Card) Portrait(source string) *Card {
	c.portrait = source

	return c
}
