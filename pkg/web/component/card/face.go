package card

import "github.com/funtimecoding/soil/pkg/web/constant"

func (c *Card) Face() *Card {
	c.treatment = constant.CardFaceClass

	return c
}
