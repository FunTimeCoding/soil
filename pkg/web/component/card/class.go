package card

func (c *Card) Class(class string) *Card {
	c.class = class

	return c
}
