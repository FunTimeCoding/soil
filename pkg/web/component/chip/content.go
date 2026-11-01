package chip

import "maragu.dev/gomponents"

func (c *Chip) Content(n ...gomponents.Node) *Chip {
	c.content = append(c.content, n...)

	return c
}
