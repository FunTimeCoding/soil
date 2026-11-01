package chip

import "maragu.dev/gomponents"

func (c *Chip) OnClick(script string) *Chip {
	c.attributes = append(c.attributes, gomponents.Attr("onclick", script))

	return c
}
