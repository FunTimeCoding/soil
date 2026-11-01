package card

import "maragu.dev/gomponents"

func (c *Card) Marker(n gomponents.Node) *Card {
	c.marker = n

	return c
}
