package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Art() *Card {
	c.treatment = constant.CardArtClass

	return c
}
