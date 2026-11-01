package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Hero() *Card {
	c.size = constant.CardHeroClass

	return c
}
