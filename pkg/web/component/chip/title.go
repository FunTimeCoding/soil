package chip

import "maragu.dev/gomponents/html"

func (c *Chip) Title(text string) *Chip {
	c.attributes = append(c.attributes, html.Title(text))

	return c
}
