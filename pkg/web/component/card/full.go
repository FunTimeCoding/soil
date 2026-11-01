package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Full() *Card {
	c.treatment = constant.CardFullClass

	return c
}
