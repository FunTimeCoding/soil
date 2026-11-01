package card

import "maragu.dev/gomponents"

func (c *Card) State(n ...gomponents.Node) *Card {
	c.state = append(c.state, n...)

	return c
}
