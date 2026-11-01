package card

import "maragu.dev/gomponents"

func (c *Card) Trailing(n ...gomponents.Node) *Card {
	c.trailing = append(c.trailing, n...)

	return c
}
