package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Compact() *Card {
	c.size = constant.CardCompactClass

	return c
}
