package card

func (c *Card) Identifier(identifier string) *Card {
	c.identifier = identifier

	return c
}
