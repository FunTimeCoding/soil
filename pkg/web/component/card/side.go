package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Side() *Card {
	c.treatment = constant.CardSideClass

	return c
}
