package card

func (c *Card) Task() *Card {
	c.task = true

	return c
}
