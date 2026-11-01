package card

import "maragu.dev/gomponents"

func (c *Card) Footer(n ...gomponents.Node) *Card {
	c.footer = append(c.footer, n...)

	return c
}
