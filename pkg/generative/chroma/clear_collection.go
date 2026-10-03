package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func (c *Client) ClearCollection(l v2.Collection) {
	c.Delete(l, v2.WithWhere(v2.NotEqString(constant.ChromaNameField, "")))
}
